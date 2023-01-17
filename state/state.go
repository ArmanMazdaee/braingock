package state

import "errors"

// Returned by IncrementPointer and DecrementPointer state method on
// moving pointer out of state acceptable range
var ErrOutOfBound = errors.New("moving to out of bound of state")

// Generic type of each state word
type StateType interface {
	uint8 | uint16 | uint32 | uint64
}

// State interface define the set of changes which can be done on a state
type State[T StateType] interface {
	// get the value which the pointer is pointing to
	PointerValue() T
	// set the value which the pointer is pointing to
	SetPointerValue(T)
	// move pointer right
	IncrementPointer() error
	// move pointer left
	DecrementPointer() error
}
