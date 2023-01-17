package state

// A fixed size state, implemented using slices for performance
type FixedSize[T StateType] struct {
	store   []T
	pointer int
}

// Return a fixed size state
func NewFixedSize[T StateType](size uint) *FixedSize[T] {
	return &FixedSize[T]{
		store:   make([]T, size),
		pointer: 0,
	}
}

func (s *FixedSize[T]) PointerValue() T {
	return s.store[s.pointer]
}

func (s *FixedSize[T]) SetPointerValue(value T) {
	s.store[s.pointer] = value
}

func (s *FixedSize[T]) IncrementPointer() error {
	if s.pointer+1 >= len(s.store) {
		return ErrOutOfBound
	}
	s.pointer = s.pointer + 1
	return nil
}

func (s *FixedSize[T]) DecrementPointer() error {
	if s.pointer <= 0 {
		return ErrOutOfBound
	}
	s.pointer = s.pointer - 1
	return nil
}
