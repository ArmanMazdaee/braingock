package state

import "testing"

func TestDynamicSizeNonLeftExpandable(t *testing.T) {
	state := NewDynamicSize[uint8](false)

	err := state.DecrementPointer()
	if err != ErrOutOfBound {
		t.Fatal("should return ErrOutOfBound on fist DecrementPointer")
	}

	err = state.IncrementPointer()
	if err != nil {
		t.Fatal("should not return error on IncrementPointer:", err)
	}

	err = state.DecrementPointer()
	if err != nil {
		t.Fatal("should not return error on second DecrementPointer:", err)
	}

	err = state.DecrementPointer()
	if err != ErrOutOfBound {
		t.Fatal("should return ErrOutOfBound on third DecrementPointer")
	}
}

func TestDynamicSizeLeftExpandable(t *testing.T) {
	state := NewDynamicSize[uint8](true)

	err := state.DecrementPointer()
	if err != nil {
		t.Fatal("should not return error on first DecrementPointer:", err)
	}
	if state.store.Len() != 2 {
		t.Fatal("state store should have len 2")
	}

	err = state.IncrementPointer()
	if err != nil {
		t.Fatal("should not return error on IncrementPointer:", err)
	}

	err = state.DecrementPointer()
	if err != nil {
		t.Fatal("should not return error on second DecrementPointer:", err)
	}
	if state.store.Len() != 2 {
		t.Fatal("state store should have len 2")
	}

	err = state.DecrementPointer()
	if err != nil {
		t.Fatal("should not return error on third DecrementPointer")
	}
	if state.store.Len() != 3 {
		t.Fatal("state store should have len 3")
	}
}

func TestDynamicSizeIncrementPointer(t *testing.T) {
	state := NewDynamicSize[uint8](false)
	err := state.IncrementPointer()
	if err != nil {
		t.Fatal("should not return error on IncrementPointer:", err)
	}
	if state.store.Len() != 2 {
		t.Fatal("state store should have len 2")
	}

	err = state.DecrementPointer()
	if err != nil {
		t.Fatal("should not return error on first DecrementPointer:", err)
	}

	err = state.IncrementPointer()
	if err != nil {
		t.Fatal("should not return error on IncrementPointer:", err)
	}

	err = state.IncrementPointer()
	if err != nil {
		t.Fatal("should not return error on IncrementPointer:", err)
	}
	if state.store.Len() != 3 {
		t.Fatal("state store should have len 3")
	}
}
