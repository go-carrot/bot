package fsm

type State struct {
	EntryAction func() error
	Transition  func(interface{}) *State
}
