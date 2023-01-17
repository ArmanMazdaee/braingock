package commands

import "github.com/ArmanMazdaee/braingock/state"

// For traking uniqe values as each call to make(chan interface{})
// would return a uniqe value
type Symbol chan interface{}

// Return a new Symbol
func NewSymbol() Symbol {
	return make(Symbol)
}

// Define the Controller interface which Commands can use to controll
// state and execution of the program
type Controller[T state.StateType] interface {
	state.State[T]
	Symbol() Symbol
	PushScope(symbol Symbol, needed, lazy bool)
	PopScope() Symbol
	Rewind()
}

// Define Command interface. HandleCommand would be called each time
// the command need to be run by the interpreter.
type Command[T state.StateType] interface {
	HandleCommand(passive bool, vm Controller[T]) error
}

// convert a function to a Command and will call the provided function
// on HandleCommand
type CommandFn[T state.StateType] func(bool, Controller[T]) error

func (fn CommandFn[T]) HandleCommand(passive bool, wm Controller[T]) error {
	return fn(passive, wm)
}
