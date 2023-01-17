package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/ArmanMazdaee/braingock"
	"github.com/ArmanMazdaee/braingock/computation"
	"github.com/ArmanMazdaee/braingock/state"
)

type orchestrationError struct {
	path  string
	index int
	err   error
}

func (e orchestrationError) Error() string {
	return fmt.Sprintf("error on source %d, %s: %v", e.index+1, e.path, e.err)
}

func (e orchestrationError) Unwrap() error {
	return e.err
}

func computeAsync[T state.StateType](
	ctx context.Context,
	inter braingock.Interpreter[T],
	path string,
	index int,
) (chan<- interface{}, <-chan orchestrationError) {
	startCh := make(chan interface{})
	errCh := make(chan orchestrationError, 1)
	go func() {
		defer close(errCh)

		source, err := openFile(path)
		if err != nil {
			errCh <- orchestrationError{path, index, err}
			return
		}

		comp := inter.ComputeWithContext(ctx, source)
		select {
		case <-ctx.Done():
			errCh <- orchestrationError{path, index, ctx.Err()}
			return
		case <-startCh:
		}

		for {
			err := comp.Step()
			if err == nil || errors.As(err, &computation.UnknownCommandError{}) {
				continue
			}
			if err == io.EOF {
				return
			}
			errCh <- orchestrationError{path, index, err}
			return
		}
	}()
	return startCh, errCh
}

// Will run the sources in order. Also is responsiable for preloading the sources
func orchestrateComputation[T state.StateType](
	ctx context.Context,
	inter braingock.Interpreter[T],
	paths []string,
) <-chan error {
	outCh := make(chan error, 1)
	go func() {
		defer close(outCh)

		var prevErrCh <-chan orchestrationError
		for index, path := range paths {
			startCh, errCh := computeAsync(ctx, inter, path, index)
			if prevErrCh != nil {
				for err := range prevErrCh {
					outCh <- err
				}
			}
			prevErrCh = errCh
			close(startCh)
		}
		if prevErrCh != nil {
			for err := range prevErrCh {
				outCh <- err
			}
		}
	}()
	return outCh
}
