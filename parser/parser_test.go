package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	l "lex_an/lexer"
)

func TestInitialState(t *testing.T) {
	equation := ")a+b"
	lexer := &l.Lexer{}
	tokens := lexer.Check(equation)
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	context.Validate(tokens)
	assert.Equal(t, 1, len(context.Errors))
}