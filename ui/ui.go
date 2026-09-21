package ui

import (
	"strings"
)

type UI struct {
	Input string
}

func (ui *UI) GetInput(eq string) {
	ui.Input = strings.ReplaceAll(eq, " ","")
}