# pomo-cli

A Pomodoro Timer CLI tool. This tool can help you stay focussed for periods of time and improve flow during tasks. It's based on the [Pomodoro Technique](https://en.wikipedia.org/wiki/Pomodoro_Technique), and while there are many alternative apps out there in the browser or as fully fledged desktop apps, for those of us who spend much of thier days in the terminal, I feel as though this is an extra distraction and interrupts flow by forcing you out of context. This tool has been developed as a CLI tool to address these issues.

## Demo

![til](./demo/demo.gif)

## Installation Instructions

### MacOS through brew

```
brew tap codeanish/pomo-cli
brew install pomo-cli
```

## Usage instructions

Launch the pomo app in your terminal of choice using `pomo`. Select one of the timer options using the up and down arrows, alternatively VIM bindings j and k also work for navigating the options. Once you've selected an option, hit `Enter` which will start the timer. Once the timer finishes, the program will exit.

Custom times can be run by just entering a time in minutes using the number keys on your keyboard and selecting the custom option.