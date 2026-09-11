package state_machine

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	equation := ")a+b"
	errors := make(map[int]string)
	context := &EquationParser{Equation: equation, Errors: errors}
	state := &InitValidationState{Context: context}
	context.SetState(state)
	context.Validate(")a+b$&")
	assert.Equal(t, 3, len(context.Errors))
}