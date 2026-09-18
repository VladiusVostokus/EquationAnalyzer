package parser

import l "lex_an/lexer"

type StateFactory func(context *EquationParser) State

var Transitions map[l.TokenType]StateFactory = map[l.TokenType]StateFactory {
    l.OP_PAREN: NewOpParenthes,
    l.CL_PAREN: NewClParenthes,
    l.OPERATOR: NewOperator,
    l.NUMBER: NewNumber,
    l.IDENTIFIER: NewIdentifier,
	l.INVALID: NewInvalid,
}

func NewOpParenthes(context *EquationParser) State {
    return &OpParenthes {Context: context}
}

func NewClParenthes(context *EquationParser) State {
    return &ClParenthes {Context: context}
}

func NewOperator(context *EquationParser) State {
    return &Operator {Context: context}
}

func NewNumber(context *EquationParser) State {
    return &Number {Context: context}
}

func NewIdentifier(context *EquationParser) State {
    return &Identifier {Context: context}
}

func NewInvalid(context *EquationParser) State {
    return &Invalid {Context: context}
}