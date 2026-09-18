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

type Operator struct {
    Context *EquationParser
}

func (s *Operator) Validate (eq []l.Token, curPos int) {}


type Number struct {
    Context *EquationParser
}

func (s *Number) Validate (eq []l.Token, curPos int) {}


type Identifier struct {
    Context *EquationParser
}

func (s *Identifier) Validate (eq []l.Token, curPos int) {}


type OpParenthes struct {
    Context *EquationParser
}

func (s *OpParenthes) Validate (eq []l.Token, curPos int) {}


type ClParenthes struct {
    Context *EquationParser
}

func (s *ClParenthes) Validate (eq []l.Token, curPos int) {}


type Invalid struct {
    Context *EquationParser
}

func (s *Invalid) Validate (eq []l.Token, curPos int) {}

