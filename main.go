package main

import (
	"fmt"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
)

func main() {
	opts := []options{
		{"Focus", 25},
		{"Focus", 15},
		{"Break", 5},
		{"Break", 15},
	}
	initialModel := model{0, false, 1000, 0, 0, false, opts}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return tick()
}

// Main update function.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if !m.Chosen {
		return updateChoices(msg, m)
	}
	return updateCountdown(msg, m)
}

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n  Time's up! Wanna go again?\n\n"
	}
	if !m.Chosen {
		s = optionsView(m)
	} else {
		s = countdownView(m)
	}
	return indent.String("\n"+s+"\n\n", 2)
}

// Sub-update functions

// Update loop for the first view where you're choosing a task.
func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
			m.Ticks = m.Options[m.Choice].Time * 60
			return m, frame()
		}

	case tickMsg:
		if m.Ticks == 0 {
			m.Quitting = true
			return m, tea.Quit
		}
		m.Ticks--
		return m, tick()
	}

	return m, nil
}

func updateCountdown(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		if m.Ticks == 0 {
			m.Quitting = true
			return m, tea.Quit
		}
		m.Ticks--
		m.Progress = (1 - float64(m.Ticks)/float64(m.Options[m.Choice].Time*60))
		return m, tick()

	}
	return m, nil
}

// Sub-view functions

func optionsView(m model) string {
	c := m.Choice

	tpl := "Time to focus?\n\n"
	for i := 0; i < len(m.Options); i++ {
		tpl += fmt.Sprintf("%s\n", checkbox(fmt.Sprintf("%s Time - %d mins", m.Options[i].Type, m.Options[i].Time), c == i))
	}
	tpl += "\n" + subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")
	return tpl
}

func countdownView(m model) string {
	var msg string

	msg = fmt.Sprintf("Focus time %s mins\n", keyword(strconv.Itoa(m.Options[m.Choice].Time)))

	var label string
	if m.Ticks > 60 {
		minutes := m.Ticks / 60
		seconds := m.Ticks % 60
		label = fmt.Sprintf("Time remaining %s minutes %s seconds...\n\n", colorFg(strconv.Itoa(minutes), "79"), colorFg(strconv.Itoa(seconds), "79"))
	} else {
		label = fmt.Sprintf("Time remaining %s seconds...\n", colorFg(strconv.Itoa(m.Ticks), "79"))
	}

	return msg + "\n\n" + label + "\n" + progressbar(m.Progress) + "%"
}
