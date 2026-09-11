package state_machine


type EquationParser struct {
    State State
    Equation string
    Errors map[int]string
}

func (ep *EquationParser) SetState(s State) {
    ep.State = s
}

func (ep *EquationParser) Validate(eq string) {
    ep.State.Validate(eq)
}