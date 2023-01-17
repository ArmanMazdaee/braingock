package commands

import (
	"reflect"
	"testing"
)

func TestStartLoop(t *testing.T) {
	controller := newMockedController[uint8]()
	command := NewStartLoop[uint8]()
	command.HandleCommand(true, controller)
	if !reflect.DeepEqual(controller.calls, []string{"PushScope"}) {
		t.Fatal("calls to controller are incorrect in passive mode")
	}
	if !reflect.DeepEqual(controller.pushScopePointerPayloads, []pushScopePayload{{loopSymbol, false, true}}) {
		t.Fatal("value pass to PushScope is incorrect")
	}

	controller = newMockedController[uint8]()
	command.HandleCommand(false, controller)
	if !reflect.DeepEqual(controller.calls, []string{"PointerValue", "PushScope"}) {
		t.Fatal("calls to controller are incorrect in active mode")
	}
	if !reflect.DeepEqual(controller.pushScopePointerPayloads, []pushScopePayload{{loopSymbol, false, true}}) {
		t.Fatal("value pass to PushScope is incorrect")
	}

	controller = newMockedController[uint8]()
	controller.value = 1
	command.HandleCommand(false, controller)
	if !reflect.DeepEqual(controller.calls, []string{"PointerValue", "PushScope"}) {
		t.Fatal("calls to controller are incorrect in active mode")
	}
	if !reflect.DeepEqual(controller.pushScopePointerPayloads, []pushScopePayload{{loopSymbol, true, false}}) {
		t.Fatal("value pass to PushScope is incorrect")
	}
}

func TestEndLoop(t *testing.T) {
	controller := newMockedController[uint8]()
	command := NewEndLoop[uint8]()
	err := command.HandleCommand(false, controller)
	if err != ErrNotLoopScope {
		t.Fatal("expect command to return ErrNotLoopScope")
	}
	if !reflect.DeepEqual(controller.calls, []string{"Symbol"}) {
		t.Fatal("calls to controller are incorrect")
	}

	controller = newMockedController[uint8]()
	controller.symbol = loopSymbol
	command.HandleCommand(true, controller)
	if !reflect.DeepEqual(controller.calls, []string{"Symbol", "PopScope"}) {
		t.Fatal("calls to controller are incorrect in passive mode")
	}

	controller = newMockedController[uint8]()
	controller.symbol = loopSymbol
	controller.value = 0
	command.HandleCommand(false, controller)
	if !reflect.DeepEqual(controller.calls, []string{"Symbol", "PointerValue", "PopScope"}) {
		t.Fatal("calls to controller are incorrect in active-false mode")
	}

	controller = newMockedController[uint8]()
	controller.symbol = loopSymbol
	controller.value = 1
	command.HandleCommand(false, controller)
	if !reflect.DeepEqual(controller.calls, []string{"Symbol", "PointerValue", "Rewind"}) {
		t.Fatal("calls to controller are incorrect in active-true mode")
	}
}
