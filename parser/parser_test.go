package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	l "lex_an/lexer"
	"reflect"
)

func TestStateTransition(t *testing.T) {
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	assert.Equal(t, reflect.TypeOf(context.State), reflect.TypeOf(&Start{}))
	state2 := &Operator{Context: context}
	context.SetState(state2)
	assert.Equal(t, reflect.TypeOf(context.State), reflect.TypeOf(&Operator{}))
}

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

func TestStartState3(t *testing.T) {
	equation := "/a+b"
	lexer := &l.Lexer{}
	tokens := lexer.Check(equation)
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	context.Validate(tokens)
	assert.Equal(t, 1, len(context.Errors))
}

func TestStartState4(t *testing.T) {
	equation := "-a+b"
	lexer := &l.Lexer{}
	tokens := lexer.Check(equation)
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	context.Validate(tokens)
	assert.Equal(t, 0, len(context.Errors))
}

func TestStartParCount(t *testing.T) {
	equation := "(a+b)"
	lexer := &l.Lexer{}
	tokens := lexer.Check(equation)
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	context.Validate(tokens)
	assert.Equal(t, context.OpenedParCount, 1)
}