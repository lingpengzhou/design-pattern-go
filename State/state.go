package State

import "log"

type State interface {
	doAction()
}

type Context struct {
	state State
}

func SetState(state State) *Context {
	return &Context{
		state,
	}
}

type StartState struct {
}

func (startState *StartState) doAction() {
	log.Println("please Open The State")
}

type StopState struct {
}

func (stopState *StopState) doAction() {
	log.Println("please Close The State")
}
