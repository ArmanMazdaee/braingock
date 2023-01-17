package computation

import (
	"fmt"
	"io"
)

// Return by Step if there are open scopes in the end of source code
var ErrDanglingScope = fmt.Errorf("scopes stack is not empty: %w", io.EOF)

// Return by Step if there an error returned by a command
type CommandError struct {
	err    error
	line   int
	column int
	char   rune
}

func (e CommandError) Error() string {
	return fmt.Sprintf(
		"command for '%c' at line %d and column %d returns error: %v",
		e.char, e.line, e.column, e.err,
	)
}

func (e CommandError) Unwrap() error {
	return e.err
}

// Return by Step if source contain an unknown command
type UnknownCommandError struct {
	line   int
	column int
	char   rune
}

func (e UnknownCommandError) Error() string {
	return fmt.Sprintf(
		"command for %c at line %d and column %d is unknown",
		e.char, e.line, e.column,
	)
}
