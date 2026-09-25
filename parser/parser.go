package parser

import (
    l "lex_an/lexer"
)

type EquationParser struct {
    State State
    Errors map[int]string
    OpenedParCount, ClosedParCount, NowOpened, lastOpPos int
}

func (ep *EquationParser) SetState(s State) {
    ep.State = s
}

func (ep *EquationParser) CountParenthes(token l.Token) {
    if token.Type == l.OP_PAREN { 
        ep.OpenedParCount++ 
        ep.NowOpened++
        ep.lastOpPos = token.Pos
    }
    if token.Type == l.CL_PAREN { 
        ep.ClosedParCount++
        ep.NowOpened--
    }
    if ep.NowOpened < 0 {
        ep.Errors[token.Pos] = "Unclosed parenthes inside equation"
        ep.NowOpened = 0
    }
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
    ep.NowOpened = 0
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
    if ep.OpenedParCount == ep.ClosedParCount && ep.NowOpened > 0 {
        ep.Errors[ep.lastOpPos] += "Equation has unclosed parenthes"
    }
}