package parser

import (
    l "lex_an/lexer"
)


type EquationParser struct {
    State State
    Errors map[int]string
    openedParCount, closedParCount int
}

func (ep *EquationParser) SetState(s State) {
    ep.State = s
}

func (ep *EquationParser) Validate(tokens []l.Token) {
    ep.Errors = make(map[int]string)
    ep.State.Validate(tokens, 0)
}