package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInitialState(t *testing.T) {
	equation := "+-*/"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "*", result[2].Value)
}