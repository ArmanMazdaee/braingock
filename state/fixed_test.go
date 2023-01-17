package state

import "testing"

func TestFixedSizeBound(t *testing.T) {
	sizes := []uint{1, 5, 10}
	for _, size := range sizes {
		state := NewFixedSize[uint8](size)
		err := state.DecrementPointer()
		if err != ErrOutOfBound {
			t.Fatal("expected ErrOutOfBound on DecrementPointer")
		}
		for i := uint(1); i < size; i++ {
			err := state.IncrementPointer()
			if err != nil {
				t.Fatalf("did not expected error on %d IncrementPointer: %v", i, err)
			}
		}
		err = state.IncrementPointer()
		if err != ErrOutOfBound {
			t.Fatalf("expected ErrOutOfBound on %d IncreatPointer", size)
		}
		for i := uint(1); i < size; i++ {
			err := state.DecrementPointer()
			if err != nil {
				t.Fatalf("did not expected error on %d DecrementPointer: %v", i, err)
			}
		}
		err = state.DecrementPointer()
		if err != ErrOutOfBound {
			t.Fatalf("expected ErrOutOfBound on %d DecrementPointer", size)
		}
	}
}
