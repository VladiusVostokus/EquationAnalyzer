package ui

import (
	"strings"
	l "lex_an/lexer"
	p "lex_an/parser"
)

type UI struct {
	Input string
	lexer *l.Lexer
	parser *p.EquationParser
}

func (ui *UI) GetInput(eq string) {
	ui.Input = strings.ReplaceAll(eq, " ","")
}

func (ui *UI) MakeCheck() {
	tokens := ui.lexer.Check(ui.Input)
	ui.parser.SetState(&p.Start{Context: ui.parser})
	ui.parser.Validate(tokens)
}

func (ui *UI) FormOutputErrors() string {
	lexErrors := ui.lexer.Errors
	syntxErrors := ui.parser.Errors
	output := ""
	errMarks = [len(eq)]string{}
	for i := 0; i < len(eq) - 1; i++ {
		errMarks[i] = " "
	}
	errNumber := []string{}
	for i, v := range lexErrors {
		errNumber[i] = v
	}
	for i, v := range syntxErrors {
		if errNumber[i] != " "{
			errNumber[i] += " ," + v
		}
		errNumber[i] = v
	}
	for i := 0; i < len(eq) - 1; i++ {
		if errNumber[i] != " "{
			errMarks[i] = "^"
		}
	}
	output += ui.Input + "\n"
	output += strings.Joint(errMarks,"") + "\n\n"
	for i, v := range(errNumber) {
		output += string(i) + v + "\n"
	}
	return output
}

func (ui *UI) Clear() {
	clear(ui.lexer.Errors)
	clear(ui.parser.Errors)
}