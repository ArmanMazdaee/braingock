package commands

import "github.com/ArmanMazdaee/braingock/state"

type pushScopePayload struct {
	symbol Symbol
	needed bool
	lazy   bool
}

type mockedController[T state.StateType] struct {
	calls                    []string
	setPointerValuePayloads  []T
	pushScopePointerPayloads []pushScopePayload
	value                    T
	symbol                   Symbol
}

func newMockedController[T state.StateType]() *mockedController[T] {
	return &mockedController[T]{
		calls:                    make([]string, 0),
		setPointerValuePayloads:  make([]T, 0),
		pushScopePointerPayloads: make([]pushScopePayload, 0),
	}
}

func (c *mockedController[T]) PointerValue() T {
	c.calls = append(c.calls, "PointerValue")
	return c.value
}

func (c *mockedController[T]) SetPointerValue(value T) {
	c.calls = append(c.calls, "SetPointerValue")
	c.setPointerValuePayloads = append(c.setPointerValuePayloads, value)
}

func (c *mockedController[T]) IncrementPointer() error {
	c.calls = append(c.calls, "IncrementPointer")
	return nil
}

func (c *mockedController[T]) DecrementPointer() error {
	c.calls = append(c.calls, "DecrementPointer")
	return nil
}

func (c *mockedController[T]) Symbol() Symbol {
	c.calls = append(c.calls, "Symbol")
	return c.symbol
}

func (c *mockedController[T]) PushScope(symbol Symbol, needed, lazy bool) {
	c.calls = append(c.calls, "PushScope")
	c.pushScopePointerPayloads = append(c.pushScopePointerPayloads, pushScopePayload{symbol, needed, lazy})
}

func (c *mockedController[T]) PopScope() Symbol {
	c.calls = append(c.calls, "PopScope")
	return c.symbol
}

func (c *mockedController[T]) Rewind() {
	c.calls = append(c.calls, "Rewind")
}
