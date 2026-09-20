package parser

import (
    l "lex_an/lexer"
)

type EquationParser struct {
    State State
    Errors map[int]string
    OpenedParCount, ClosedParCount int
}

func (ep *EquationParser) SetState(s State) {
    ep.State = s
}

func (ep *EquationParser) Validate(tokens []l.Token) {
    ep.Errors = make(map[int]string)
    for i, _ := range(tokens) {
        ep.State.Validate(tokens, i)
    }
    lastIdx := len(tokens) - 1
    lastToken := tokens[lastIdx]
    if lastToken.Type == l.OPERATOR {
        ep.Errors[lastIdx] = "Unexpected symbol at the end of equation:" + lastToken.Value
    }
    if ep.OpenedParCount != ep.ClosedParCount {
        ep.Errors[len(tokens)] = "Equation have unclosed parenthes"
    }
}