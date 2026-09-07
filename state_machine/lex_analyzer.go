package state_machine

import (
    "strings"
)

type EquationParser struct {
    state State
    equation string
}