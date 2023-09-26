package main

type model struct {
	Choice   int
	Chosen   bool
	Ticks    int
	Frames   int
	Progress float64
	Quitting bool
	Options  []options
}

type options struct {
	Type string
	Time int
}

type (
	tickMsg  struct{}
	frameMsg struct{}
)
