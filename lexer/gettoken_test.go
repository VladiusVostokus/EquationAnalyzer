package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	assert.Equal(t, OPERATOR, GetToken["+"])
	assert.Equal(t, OPERATOR, GetToken["-"])
	assert.Equal(t, OPERATOR, GetToken["*"])
	assert.Equal(t, OPERATOR, GetToken["/"])
	assert.Equal(t, OP_PAREN, GetToken["("])
	assert.Equal(t, CL_PAREN, GetToken[")"])
	assert.Equal(t, DOT, GetToken["."])
}