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

func TestSimpleEquation(t *testing.T) {
	equation := "12+8/2"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 5, len(result))
}

func TestSimpleEquationWithParents(t *testing.T) {
	equation := "(12+8)/2"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 7, len(result))
}

func TestNumberWithDots(t *testing.T) {
	equation := "12.25"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 1, len(result))
}

func TestEquationWithDots(t *testing.T) {
	equation := "(12.25+11.3)*4"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 7, len(result))
	assert.Equal(t, result[1].Type, NUMBER)
}

func TestIdentifier(t *testing.T) {
	equation := "a"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, result[0].Type, IDENTIFIER)
}

func TestComplexEquation(t *testing.T) {
	equation := "-ab*(12.25+11.3)*4/6.7"
	l := &Lexer{}
	result := l.Check(equation)
	assert.Equal(t, 13, len(result))
}