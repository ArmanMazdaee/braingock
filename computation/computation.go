package computation

import (
	"context"
	"io"

	"github.com/ArmanMazdaee/braingock/commands"
	"github.com/ArmanMazdaee/braingock/state"
)

// Represent the computaion of a source code against specified state and commands
type Computation[T state.StateType] struct {
	ctx      context.Context
	cancel   context.CancelFunc
	state    state.State[T]
	commands map[rune]commands.Command[T]
	tape     *tape
}

// Return a new computaion. The reader get closed on cancelation or end of execution
func New[T state.StateType](
	ctx context.Context,
	state state.State[T],
	commandsMap map[rune]commands.Command[T],
	source io.ReadCloser,
) *Computation[T] {
	innerCtx, cancel := context.WithCancel(ctx)
	if commandsMap == nil {
		commandsMap = make(map[rune]commands.Command[T])
	}
	tape := newTape(ctx, source)
	return &Computation[T]{innerCtx, cancel, state, commandsMap, tape}
}

// Move the computaion one step forward.
func (c *Computation[T]) Step() error {
	if err := c.ctx.Err(); err != nil {
		return err
	}

	token, passive, err := c.tape.nextToken()
	if err != nil {
		return err
	}

	command, ok := c.commands[token.char]
	if !ok || command == nil {
		c.tape.dropToken()
		return UnknownCommandError(token)
	}

	controller := struct {
		state.State[T]
		*tape
	}{c.state, c.tape}
	if err := command.HandleCommand(passive, controller); err != nil {
		return CommandError{err, token.line, token.column, token.char}
	}
	return nil
}

// Cancel the computation
func (c *Computation[T]) Cancel() {
	if err := c.ctx.Err(); err == nil {
		c.cancel()
	}
}
