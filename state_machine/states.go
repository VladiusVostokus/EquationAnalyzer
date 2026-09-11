package state_machine

import (
    "strings"
    "golang.org/x/exp/slices"
)

var forbidenStarEndSymbols = []string{"+","-","*","/","&","$","_","@","#",".","!","?","^"}
var forbidenSymbols = []string{"&","$","_","@","#",".","!","?","^"}

type State interface {
    Validate(eq string)
}

type InitValidationState struct {
    Context *EquationParser
}

func (ivs *InitValidationState) Validate (eq string) {
    eqStart := string(eq[0])
    if strings.HasPrefix(eq, ")") || slices.Contains(forbidenStarEndSymbols, eqStart) {
        ivs.Context.Errors[0] = "Wrong symbol to start"
    }
    endIdx := len(eq) - 1
    eqEnd := string(eq[endIdx])
    if strings.HasSuffix(eq, "(") || slices.Contains(forbidenStarEndSymbols, eqEnd) {
        ivs.Context.Errors[endIdx] = "Wrong symbol to end"
    }
}