package commands

import (
	"errors"

	"github.com/ArmanMazdaee/braingock/state"
)

// Return from endLoop if last opened scope is not a loop scope
var ErrNotLoopScope = errors.New("top scope is not a loop scope")

// Used for tracking loop scopes
var loopSymbol = NewSymbol()

// If pointer value is 0 skip till next endLoop, else move to next command
func NewStartLoop[T state.StateType]() Command[T] {
	return CommandFn[T](func(passive bool, controller Controller[T]) error {
		truth := !passive && controller.PointerValue() != 0
		controller.PushScope(loopSymbol, truth, !truth)
		return nil
	})
}

// If pointer value is not 0 return to the last openLoop, else move to next command
func NewEndLoop[T state.StateType]() Command[T] {
	return CommandFn[T](func(passive bool, controller Controller[T]) error {
		if controller.Symbol() != loopSymbol {
			return ErrNotLoopScope
		}
		if passive || controller.PointerValue() == 0 {
			controller.PopScope()
			return nil
		}
		controller.Rewind()
		return nil
	})
}
