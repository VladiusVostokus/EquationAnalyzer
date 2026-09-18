package parser

import (
    l "lex_an/lexer"
    "golang.org/x/exp/slices"
)

type State interface {
    Validate(eq []l.Token, pos int)
}

type Start struct {
    Context *EquationParser
}

var allowedStartTokens = []l.TokenType{l.OP_PAREN, l.NUMBER, l.IDENTIFIER}

func (s *Start) Validate (eq []l.Token, curPos int) {
    firstToken := eq[curPos]
    if !slices.Contains(allowedStartTokens, firstToken.Type) && firstToken.Value != "-"{
        s.Context.Errors[curPos] = "Unexpected symbol at start of equation:" + firstToken.Value
    } else {
        if firstToken.Type == l.OP_PAREN {
            s.Context.OpenedParCount++
        }
    }
}