package parser

import (
    l "lex_an/lexer"
    "golang.org/x/exp/slices"
)

var allowedStartTokens = []l.TokenType{l.OP_PAREN, l.NUMBER, l.IDENTIFIER, l.INVALID}
var allowedOperatorTokens = allowedStartTokens
var allowedNumberTokens = []l.TokenType{l.IDENTIFIER, l.OPERATOR, l.CL_PAREN, l.INVALID}
var allowedIdentifierTokens = []l.TokenType{l.IDENTIFIER, l.CL_PAREN, l.OPERATOR}
var allowedOpParenTokens = []l.TokenType{l.IDENTIFIER, l.NUMBER, l.OPERATOR, l.INVALID}
var allowedClParenTokens = []l.TokenType{l.OPERATOR, l.INVALID}

type State interface {
    Validate(eq []l.Token, pos int)
}

type Start struct {
    Context *EquationParser
}

func (s *Start) Validate (eq []l.Token, curPos int) {
    firstToken := eq[curPos]
    if !slices.Contains(allowedStartTokens, firstToken.Type) && firstToken.Value != "-"{
        s.Context.Errors[firstToken.Pos] = "Unexpected symbol at start of equation:" + firstToken.Value
        s.Context.SetState(&Invalid{s.Context})
    } else {
        state, ok := Transitions[firstToken.Type]
        if ok {
            if firstToken.Type == l.OP_PAREN { s.Context.OpenedParCount++ }
            if firstToken.Type == l.CL_PAREN { s.Context.ClosedParCount++ }
            s.Context.SetState(state(s.Context))
        }
    }
}

type Operator struct {
    Context *EquationParser
}

func (s *Operator) Validate (eq []l.Token, curPos int) {
    token := eq[curPos]
    if !slices.Contains(allowedOperatorTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol at start of equation:" + token.Value
        s.Context.SetState(&Invalid{s.Context})
    } else {
        state, ok := Transitions[token.Type]
        if ok {
            if token.Type == l.OP_PAREN { s.Context.OpenedParCount++ }
            if token.Type == l.CL_PAREN { s.Context.ClosedParCount++ }
            s.Context.SetState(state(s.Context))
        }
    }
}


type Number struct {
    Context *EquationParser
}

func (s *Number) Validate (eq []l.Token, curPos int) {
    token := eq[curPos]
    if !slices.Contains(allowedNumberTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol at start of equation:" + token.Value
        s.Context.SetState(&Invalid{s.Context})
    } else {
        state, ok := Transitions[token.Type]
        if ok {
            if token.Type == l.OP_PAREN { s.Context.OpenedParCount++ }
            if token.Type == l.CL_PAREN { s.Context.ClosedParCount++ }
            s.Context.SetState(state(s.Context))
        }
    }
}


type Identifier struct {
    Context *EquationParser
}

func (s *Identifier) Validate (eq []l.Token, curPos int) {
    token := eq[curPos]
    if !slices.Contains(allowedIdentifierTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol at start of equation:" + token.Value
        s.Context.SetState(&Invalid{s.Context})
    } else {
        state, ok := Transitions[token.Type]
        if ok {
            if token.Type == l.OP_PAREN { s.Context.OpenedParCount++ }
            if token.Type == l.CL_PAREN { s.Context.ClosedParCount++ }
            s.Context.SetState(state(s.Context))
        }
    }
}


type OpParenthes struct {
    Context *EquationParser
}

func (s *OpParenthes) Validate (eq []l.Token, curPos int) {
    token := eq[curPos]
    if !slices.Contains(allowedOpParenTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol at start of equation:" + token.Value
        s.Context.SetState(&Invalid{s.Context})
    } else {
        state, ok := Transitions[token.Type]
        if ok {
            if token.Type == l.OP_PAREN { s.Context.OpenedParCount++ }
            if token.Type == l.CL_PAREN { s.Context.ClosedParCount++ }
            s.Context.SetState(state(s.Context))
        }
    }
}


type ClParenthes struct {
    Context *EquationParser
}

func (s *ClParenthes) Validate (eq []l.Token, curPos int) {
    token := eq[curPos]
    if !slices.Contains(allowedClParenTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol at start of equation:" + token.Value
        s.Context.SetState(&Invalid{s.Context})
    } else {
        state, ok := Transitions[token.Type]
        if ok {
            if token.Type == l.OP_PAREN { s.Context.OpenedParCount++ }
            if token.Type == l.CL_PAREN { s.Context.ClosedParCount++ }
            s.Context.SetState(state(s.Context))
        }
    }
}


type Invalid struct {
    Context *EquationParser
}

func (s *Invalid) Validate (eq []l.Token, curPos int) {
    token := eq[curPos]
    state, ok := Transitions[token.Type]
    if ok {
        if token.Type == l.OP_PAREN { s.Context.OpenedParCount++ }
        if token.Type == l.CL_PAREN { s.Context.ClosedParCount++ }
        s.Context.SetState(state(s.Context))
    }
}

