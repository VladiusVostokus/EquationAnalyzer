package ui

import (
	"testing"
	"github.com/stretchr/testify/assert"
	l "lex_an/lexer"
	p "lex_an/parser"
)

func TestStateFactory(t *testing.T) {
	ui := &UI{}
	equation := " a +   b/3 * 5 "
	ui.GetInput(equation)
	assert.Equal(t,"a+b/3*5", ui.Input)
}

func TestClear(t *testing.T) {
	lexer := &l.Lexer{Errors: map[int]string{0: "Some error"}}
	parser := &p.EquationParser{Errors: map[int]string{0: "Some error"}}
	ui := &UI{Lexer: lexer, Parser: parser}
	ui.Clear()
	assert.Equal(t, 0, len(lexer.Errors))
	assert.Equal(t, 0, len(parser.Errors))
}