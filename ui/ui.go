package ui

import (
	"strings"
	l "lex_an/lexer"
	p "lex_an/parser"
	"fmt"
	"sort"
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
	if len(Lexerrors) == 0 && len(syntxErrors) == 0 {
		output += "Equation has no errors\n"
		return output 
	}
	errMarks := make([]string, lens)
	for i := range(errMarks) {
		errMarks[i] = " "
	}
	errNumber := make(map[int]string)
	for i, v := range Lexerrors {
		errNumber[i] = v
	}
	for i, v := range syntxErrors {
		_, ok := errNumber[i]
		if ok {
			errNumber[i] += "," + v
		} else {
			errNumber[i] = v
		}
	}
	for i := 0; i < lens; i++ {
		_, ok := errNumber[i]
		if ok {
			errMarks[i] = "^"
		}
	}
	output += ui.Input + "\n"
	output += strings.Join(errMarks,"") + "\n\n"
	output += "Error positions (starting from 0):\n"
	keys := make([]int, 0, len(errNumber))
	for k := range(errNumber) {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, v := range(keys) {
		output += fmt.Sprintf("%d. Pos: %d, Desc: %s\n",i ,v, errNumber[v])
	}
	return output
}

func (ui *UI) Clear() {
	clear(ui.Lexer.Errors)
	clear(ui.Parser.Errors)
}