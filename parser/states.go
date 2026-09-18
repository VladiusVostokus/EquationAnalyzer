package parser

import (
    l "lex_an/lexer"
)

//var forbidenStarEndSymbols = []string{"+","-","*","/","&","$","_","@","#",".","!","?","^"}
//var forbidenSymbols = []string{"&","$","_","@","#",".","!","?","^","%"}
//var arithmeticSymbols = []string{"+","-","*","/"}


type State interface {
    Validate(eq []l.Token, pos int)
}

type Start struct {
    Context *EquationParser
}

func (ivs *Start) Validate (eq []l.Token, curPos int) {
    firstToken := eq[curPos]
    if firstToken.Type == l.CL_PAREN {
        ivs.Context.Errors[curPos] = "Unexpected ) at start of equation"
    }
    if firstToken.Type == l.INVALID {
        ivs.Context.Errors[curPos] = "Unexpected symbol at start of equation:" + firstToken.Value
    }
    if firstToken.Type == l.OPERATOR {
        if firstToken.Value != "-"{
            ivs.Context.Errors[curPos] = "Unexpected operator at start of equation:" + firstToken.Value
        }
    }
}