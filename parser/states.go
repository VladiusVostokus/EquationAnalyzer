package parser

import (
    l "lex_an/lexer"
    "golang.org/x/exp/slices"
)

var allowedStartTokens = []l.TokenType{l.OP_PAREN, l.NUMBER, l.IDENTIFIER, l.INVALID}
var allowedOperatorTokens = allowedStartTokens
var allowedNumberTokens = []l.TokenType{l.IDENTIFIER, l.OPERATOR, l.CL_PAREN, l.OP_PAREN, l.INVALID}
var allowedIdentifierTokens = []l.TokenType{l.IDENTIFIER, l.CL_PAREN,  l.OPERATOR, l.INVALID}
var allowedOpParenTokens = []l.TokenType{l.IDENTIFIER, l.NUMBER, l.INVALID, l.OP_PAREN}
var allowedClParenTokens = []l.TokenType{l.OPERATOR, l.CL_PAREN, l.INVALID}

type State interface {
    Validate(token l.Token)
}

type Start struct {
    Context *EquationParser
}

func (s *Start) Validate (token l.Token) {
    if !slices.Contains(allowedStartTokens, token.Type) && token.Value != "-"{
        s.Context.Errors[token.Pos] = "Unexpected symbol at start of equation: " + token.Value
    } 
}

type Operator struct {
    Context *EquationParser
}

func (s *Operator) Validate (token l.Token) {
    if !slices.Contains(allowedOperatorTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol: " + token.Value
    } 
}


type Number struct {
    Context *EquationParser
}

func (s *Number) Validate (token l.Token) {
    if !slices.Contains(allowedNumberTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol: " + token.Value
    }
}


type Identifier struct {
    Context *EquationParser
}

func (s *Identifier) Validate (token l.Token) {
    if !slices.Contains(allowedIdentifierTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol: " + token.Value
    }
}


type OpParenthes struct {
    Context *EquationParser
}

func (s *OpParenthes) Validate (token l.Token) {
    if !slices.Contains(allowedOpParenTokens, token.Type) && token.Value != "-"  {
        s.Context.Errors[token.Pos] = "Unexpected symbol: " + token.Value
    }
}


type ClParenthes struct {
    Context *EquationParser
}

func (s *ClParenthes) Validate (token l.Token) {
    if !slices.Contains(allowedClParenTokens, token.Type) {
        s.Context.Errors[token.Pos] = "Unexpected symbol: " + token.Value
    }
}


type Invalid struct {
    Context *EquationParser
}

func (s *Invalid) Validate (token l.Token) {}

