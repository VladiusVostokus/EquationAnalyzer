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