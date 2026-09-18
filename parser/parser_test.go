package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	l "lex_an/lexer"
)

func TestStartState(t *testing.T) {
	equation := ")a+b"
	lexer := &l.Lexer{}
	tokens := lexer.Check(equation)
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	context.Validate(tokens)
	assert.Equal(t, 1, len(context.Errors))
}

func TestStartState2(t *testing.T) {
	equation := "#a+b"
	lexer := &l.Lexer{}
	tokens := lexer.Check(equation)
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	context.Validate(tokens)
	assert.Equal(t, 1, len(context.Errors))
}