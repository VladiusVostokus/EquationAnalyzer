package state_machine

import (
    "strings"
)

type EquationParser struct {
    state State
    equation string
    errors map[int]string
}

func (ep *EquationParser) SetState(s State) {
    ep.state = s
}