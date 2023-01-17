package state

import "container/list"

// A dynamic size state, implemented using linked list for flexibility.
type DynamicSize[T StateType] struct {
	leftExpandable bool
	store          *list.List
	pointer        *list.Element
}

// Return a dynamic size state. if leftExpandable is false, would return
// ErrOutOfBound if trying to move pointer left on the first cell
func NewDynamicSize[T StateType](leftExpandable bool) *DynamicSize[T] {
	store := list.New()
	var zero T
	pointer := store.PushFront(zero)
	return &DynamicSize[T]{leftExpandable, store, pointer}
}

func (s *DynamicSize[T]) PointerValue() T {
	return s.pointer.Value.(T)
}

func (s *DynamicSize[T]) SetPointerValue(value T) {
	s.pointer.Value = value
}

func (s *DynamicSize[T]) IncrementPointer() error {
	if next := s.pointer.Next(); next != nil {
		s.pointer = next
		return nil
	}

	var zero T
	s.pointer = s.store.PushBack(zero)
	return nil
}

func (s *DynamicSize[T]) DecrementPointer() error {
	if prev := s.pointer.Prev(); prev != nil {
		s.pointer = prev
		return nil
	}

	if !s.leftExpandable {
		return ErrOutOfBound
	}

	var zero T
	s.pointer = s.store.PushFront(zero)
	return nil
}
