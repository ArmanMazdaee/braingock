package commands

import "github.com/ArmanMazdaee/braingock/state"

// Increment the pointer value by one
func NewIncrementPointerValue[T state.StateType]() Command[T] {
	return CommandFn[T](func(passive bool, controller Controller[T]) error {
		if passive {
			return nil
		}
		value := controller.PointerValue()
		value += 1
		controller.SetPointerValue(value)
		return nil
	})
}

// Decrement the pointer value by one
func NewDecrementPointerValue[T state.StateType]() Command[T] {
	return CommandFn[T](func(passive bool, controller Controller[T]) error {
		if passive {
			return nil
		}
		value := controller.PointerValue()
		value -= 1
		controller.SetPointerValue(value)
		return nil

	})
}

// Move the pointer to the right
func NewIncrementPointer[T state.StateType]() Command[T] {
	return CommandFn[T](func(passive bool, controller Controller[T]) error {
		if passive {
			return nil
		}
		return controller.IncrementPointer()
	})
}

// Move the pointer to the left
func NewDecrementPointer[T state.StateType]() Command[T] {
	return CommandFn[T](func(passive bool, controller Controller[T]) error {
		if passive {
			return nil
		}
		return controller.DecrementPointer()
	})
}
