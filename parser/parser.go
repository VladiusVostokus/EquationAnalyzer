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
}