package State

import "testing"

func Test_State(t *testing.T) {
	startState := SetState(&StartState{})
	startState.state.doAction()

	stopState := SetState(&StopState{})
	stopState.state.doAction()

}
