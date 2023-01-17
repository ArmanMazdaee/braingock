package commands

import (
	"reflect"
	"testing"
)

func TestIncrementPointerValue(t *testing.T) {
	for _, value := range []uint8{0, 10, 255} {
		controller := newMockedController[uint8]()
		controller.value = value
		command := NewIncrementPointerValue[uint8]()
		command.HandleCommand(true, controller)
		if !reflect.DeepEqual(controller.calls, []string{}) {
			t.Fatal("calls to controller are incorrect in passive mode")
		}

		controller = newMockedController[uint8]()
		controller.value = value
		command.HandleCommand(false, controller)
		if !reflect.DeepEqual(controller.calls, []string{"PointerValue", "SetPointerValue"}) {
			t.Fatal("calls to controller are incorrect in active mode")
		}
		if !reflect.DeepEqual(controller.setPointerValuePayloads, []uint8{value + 1}) {
			t.Fatal("value pass to SetPointerValue is incorrect")
		}
	}
}

func TestDecrementPointerValue(t *testing.T) {
	for _, value := range []uint8{0, 10, 255} {
		controller := newMockedController[uint8]()
		controller.value = value
		command := NewDecrementPointerValue[uint8]()
		command.HandleCommand(true, controller)
		if !reflect.DeepEqual(controller.calls, []string{}) {
			t.Fatal("calls to controller are incorrect in passive mode")
		}

		controller = newMockedController[uint8]()
		controller.value = value
		command.HandleCommand(false, controller)
		if !reflect.DeepEqual(controller.calls, []string{"PointerValue", "SetPointerValue"}) {
			t.Fatal("calls to controller are incorrect in active mode")
		}
		if !reflect.DeepEqual(controller.setPointerValuePayloads, []uint8{value - 1}) {
			t.Fatal("value pass to SetPointerValue is incorrect")
		}
	}
}

func TestIncrementPointer(t *testing.T) {
	controller := newMockedController[uint8]()
	command := NewIncrementPointer[uint8]()
	command.HandleCommand(true, controller)
	if !reflect.DeepEqual(controller.calls, []string{}) {
		t.Fatal("calls to controller are incorrect in passive mode")
	}

	controller = newMockedController[uint8]()
	command.HandleCommand(false, controller)
	if !reflect.DeepEqual(controller.calls, []string{"IncrementPointer"}) {
		t.Fatal("calls to controller are incorrect in active mode")
	}
}

func TestDecrementPointer(t *testing.T) {
	controller := newMockedController[uint8]()
	command := NewDecrementPointer[uint8]()
	command.HandleCommand(true, controller)
	if !reflect.DeepEqual(controller.calls, []string{}) {
		t.Fatal("calls to controller are incorrect in passive mode")
	}

	controller = newMockedController[uint8]()
	command.HandleCommand(false, controller)
	if !reflect.DeepEqual(controller.calls, []string{"DecrementPointer"}) {
		t.Fatal("calls to controller are incorrect in active mode")
	}
}
