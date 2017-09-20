package fsm

type Emitter interface {
	Emit(interface{}) error
}
