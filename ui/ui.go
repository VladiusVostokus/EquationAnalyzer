package ui

import (
	"strings"
	l "lex_an/lexer"
	p "lex_an/parser"
)

type UI struct {
	Input string
	Lexer *l.Lexer
	Parser *p.EquationParser
}

func (ui *UI) GetInput(eq string) {
	ui.Input = strings.ReplaceAll(eq, " ","")
}

func (ui *UI) MakeCheck() {
	tokens := ui.Lexer.Check(ui.Input)
	ui.Parser.SetState(&p.Start{Context: ui.Parser})
	ui.Parser.Validate(tokens)
}

func (ui *UI) FormOutputErrors() string {
	Lexerrors := ui.Lexer.Errors
	syntxErrors := ui.Parser.Errors
	output := ""
	eq := ui.Input
	lens := len(eq)
	errMarks := []string{}
	for i := 0; i < lens - 1; i++ {
		errMarks[i] = " "
	}
	errNumber := []string{}
	for i, v := range Lexerrors {
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
	output += strings.Join(errMarks,"") + "\n\n"
	for i, v := range(errNumber) {
		output += string(rune(i)) + v + "\n"
	}
	return output
}

func (ui *UI) Clear() {
	clear(ui.Lexer.Errors)
	clear(ui.Parser.Errors)
}