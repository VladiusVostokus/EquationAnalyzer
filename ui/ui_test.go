package ui

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStateFactory(t *testing.T) {
	ui := &UI{}
	equation := " a +   b/3 * 5 "
	ui.GetInput(equation)
	assert.Equal(t,"a+b/3*5", ui.Input)
}