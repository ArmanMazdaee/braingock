package computation

import (
	"bufio"
	"context"
	"io"
	"unicode"
)

type token struct {
	line   int
	column int
	char   rune
}

type lexer struct {
	tokenCh <-chan token
	errCh   <-chan error
	err     error
}

// Lex the source in parallel and close the source on any error
func newLexer(ctx context.Context, source io.ReadCloser) *lexer {
	tokenCh := make(chan token, 1)
	errCh := make(chan error, 1)
	go func() {
		defer source.Close()

		reader := bufio.NewReader(source)
		for line := 1; ; line++ {
			text, err := reader.ReadString('\n')
			for index, char := range text {
				if unicode.IsSpace(char) {
					continue
				}
				select {
				case tokenCh <- token{line, index + 1, char}:
					continue
				case <-ctx.Done():
					errCh <- ctx.Err()
					return
				}
			}
			if err != nil {
				errCh <- err
				return
			}
		}
	}()

	return &lexer{tokenCh, errCh, nil}
}

// Return the next token from the source
func (l *lexer) lex() (token, error) {
	if err := l.err; err != nil {
		return token{}, err
	}
	select {
	case token := <-l.tokenCh:
		return token, nil
	case err := <-l.errCh:
		l.err = err
		return token{}, err
	}
}
