package commands

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/ArmanMazdaee/braingock/state"
)

func readByte[T state.StateType](
	passive bool,
	controller Controller[T],
	reader io.Reader,
	promptWriter,
	errWriter io.Writer,
) error {
	if passive {
		return nil
	}
	for {
		if promptWriter != nil {
			fmt.Fprint(promptWriter, "enter a byte: ")
		}
		var line string
		_, err := fmt.Fscanln(reader, &line)
		if err != nil {
			fmt.Fprintln(errWriter, "could not read line:", err)
			continue
		}

		line = strings.TrimSpace(line)
		value, err := strconv.ParseUint(line, 0, 8)
		if err != nil {
			fmt.Fprintf(errWriter, "%s could not converted to a byte: %v\n", line, err)
			continue
		}
		controller.SetPointerValue(T(value))
		break
	}
	return nil
}

// Read a single byte a numeric value from reader.
// If promptWriter is not nil, it will use to show the prompt message
// if user input have any error, it will be write to errWriter
func NewReadByte[T state.StateType](
	reader io.Reader,
	promptWriter,
	errWriter io.Writer,
) Command[T] {
	return CommandFn[T](func(passive bool, controller Controller[T]) error {
		return readByte(passive, controller, reader, promptWriter, errWriter)
	})
}

func writeByte[T state.StateType](
	passive bool,
	controller Controller[T],
	writer io.Writer,
	errWriter io.Writer,
) error {
	if passive {
		return nil
	}
	value := controller.PointerValue()
	if value > 255 {
		fmt.Fprintf(errWriter, "pointer value %v does not fit in a byte", value)
	}
	fmt.Fprintln(writer, byte(value))
	return nil
}

// Write a single byte a numeric value from writer.
// if user pointer value does not fit in a byte, write an error to errWriter
func NewWriteByte[T state.StateType](writer io.Writer, errWriter io.Writer) Command[T] {
	return CommandFn[T](func(passive bool, controller Controller[T]) error {
		return writeByte(passive, controller, writer, errWriter)
	})
}
