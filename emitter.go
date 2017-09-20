package fsm

// Emitter is a generic interface to output arbitrary data.
// Emit is generally called from State.EntryAction.
type Emitter interface {
	Emit(interface{}) error
}
