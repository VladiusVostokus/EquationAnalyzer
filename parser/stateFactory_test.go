package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	l "lex_an/lexer"
	"reflect"
)

func TestStateFactory(t *testing.T) {
	context := &EquationParser{}
	state := Transitions[l.NUMBER]
	context.SetState(state(context))
	assert.Equal(t, reflect.TypeOf(context.State), reflect.TypeOf(&Number{}))
	state2 := Transitions[l.OPERATOR]
	context.SetState(state2(context))
	assert.Equal(t, reflect.TypeOf(context.State), reflect.TypeOf(&Operator{}))
}

func TestTransitToInvalidState(t *testing.T) {
	equation := ")"
	lexer := &l.Lexer{}
	tokens := lexer.Check(equation)
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	context.Validate(tokens)
	assert.Equal(t, reflect.TypeOf(context.State), reflect.TypeOf(&Invalid{}))
}

func TestTransitToIdentifierState(t *testing.T) {
	equation := ")a"
	lexer := &l.Lexer{}
	tokens := lexer.Check(equation)
	context := &EquationParser{}
	state := &Start{Context: context}
	context.SetState(state)
	context.Validate(tokens)
	assert.Equal(t, reflect.TypeOf(context.State), reflect.TypeOf(&Identifier{}))
}