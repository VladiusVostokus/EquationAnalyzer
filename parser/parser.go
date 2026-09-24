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

func (ep *EquationParser) CountParenthes(token l.Token) {
    if token.Type == l.OP_PAREN { ep.OpenedParCount++ }
    if token.Type == l.CL_PAREN { ep.ClosedParCount++ }
}

func (ep *EquationParser) ChangeState(token l.Token) {
    state, ok := Transitions[token.Type]
    if ok {
        ep.SetState(state(ep))
    }
}

func (ep *EquationParser) Validate(tokens []l.Token) {
    ep.Errors = make(map[int]string)
    ep.OpenedParCount = 0
    ep.ClosedParCount = 0
    for _, token := range(tokens) {
        ep.State.Validate(token)
        ep.CountParenthes(token)
        ep.ChangeState(token)
    }
    lastIdx := len(tokens) - 1
    lastToken := tokens[lastIdx]
    if lastToken.Type == l.OPERATOR {
        ep.Errors[lastToken.Pos] = "Unexpected symbol at the end of equation: " + lastToken.Value
    }
    if ep.OpenedParCount > ep.ClosedParCount {
        ep.Errors[lastToken.Pos + 1] = "Count of opened parenthes bigger then closed"
    }
    if ep.OpenedParCount < ep.ClosedParCount {
        ep.Errors[lastToken.Pos + 1] = "Count of closed parenthes bigger then opened"
    }
}