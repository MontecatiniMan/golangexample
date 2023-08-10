package state

type State struct {
	State string `json:"state"`
}

func NewState(state string) *State {
	return &State{
		State: state,
	}
}

func (instance *State) SetState(state string) {
	instance.State = state
}

func (instance *State) GetState() string {
	return instance.State
}
