package fsm

// State represents an individual state in a larger state machine
type State struct {
	Slug        string
	EntryAction func() error
	Transition  func(interface{}) *State
}
