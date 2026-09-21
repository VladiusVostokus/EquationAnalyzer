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