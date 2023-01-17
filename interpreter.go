package braingock

import (
	"context"
	"io"

	"github.com/ArmanMazdaee/braingock/commands"
	"github.com/ArmanMazdaee/braingock/computation"
	"github.com/ArmanMazdaee/braingock/state"
)

// A pair of state and commands list which can be use to compute brainfuck programmes
type Interpreter[T state.StateType] struct {
	state    state.State[T]
	commands map[rune]commands.Command[T]
}

func NewInterpreter[T state.StateType](
	state state.State[T],
	cmds map[rune]commands.Command[T],
) Interpreter[T] {
	return Interpreter[T]{state, cmds}
}

func (i Interpreter[T]) ComputeWithContext(ctx context.Context, source io.ReadCloser) *computation.Computation[T] {
	return computation.New(ctx, i.state, i.commands, source)
}

func (i Interpreter[T]) Compute(source io.ReadCloser) *computation.Computation[T] {
	return i.ComputeWithContext(context.Background(), source)
}
