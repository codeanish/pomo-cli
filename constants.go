package main

import "github.com/muesli/termenv"

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
)

// General stuff for styling the view
var (
	term          = termenv.EnvColorProfile()
	keyword       = makeFgStyle("211")
	subtle        = makeFgStyle("241")
	progressEmpty = subtle(progressEmptyChar)
	dot           = colorFg(" • ", "236")
	// Gradient colors we'll use for the progress bar
	ramp = makeRamp("#B14FFF", "#00FFA3", progressBarWidth)
)
