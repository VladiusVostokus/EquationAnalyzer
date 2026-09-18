package parser

import (
    l "lex_an/lexer"
)

type State interface {
    Validate(eq []l.Token, pos int)
}

type Start struct {
    Context *EquationParser
}

func (s *Start) Validate (eq []l.Token, curPos int) {
    firstToken := eq[curPos]
    if firstToken.Type == l.CL_PAREN {
        s.Context.Errors[curPos] = "Unexpected ) at start of equation"
    }
    if firstToken.Type == l.INVALID {
        s.Context.Errors[curPos] = "Unexpected symbol at start of equation:" + firstToken.Value
    }
    if firstToken.Type == l.OPERATOR {
        if firstToken.Value != "-"{
            s.Context.Errors[curPos] = "Unexpected operator at start of equation:" + firstToken.Value
        }
    }
    if firstToken.Type == l.OP_PAREN {
        s.Context.OpenedParCount++
    }
}