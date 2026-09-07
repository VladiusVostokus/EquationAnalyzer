package state_machine

import (
    "strings"
)

type State interface {
    validate(string)
}