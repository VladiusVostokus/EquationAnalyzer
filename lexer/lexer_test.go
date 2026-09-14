package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOperators(t *testing.T) {
	equation := "+-*/"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "*", result[2].Value)
	assert.Equal(t, OPERATOR, result[2].Type)
}

func TestParent(t *testing.T) {
	equation := "()"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, ")", result[1].Value)
	assert.Equal(t, CL_PAREN, result[1].Type)
}

func TestNumber(t *testing.T) {
	equation1 := "1"
	equation2 := "123"

	l1 := &Lexer{}
	l2 := &Lexer{}
	result1 := l1.Check(equation1)
	result2 := l2.Check(equation2)
	assert.Equal(t, 1, len(result1))
	assert.Equal(t, 1, len(result2))

	assert.Equal(t, "123", result2[0].Value)
	assert.Equal(t, NUMBER, result2[0].Type)
}