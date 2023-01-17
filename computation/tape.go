package computation

import (
	"context"
	"errors"
	"io"

	"github.com/ArmanMazdaee/braingock/commands"
)

type scope struct {
	symbol   commands.Symbol
	start    int
	lazy     bool
	rewinder int
}

// Handle the computaion flow of the source code
type tape struct {
	lexer  *lexer
	record int
	tokens []token
	scopes []scope
}

func newTape(ctx context.Context, source io.ReadCloser) *tape {
	return &tape{
		lexer:  newLexer(ctx, source),
		record: -1,
		tokens: make([]token, 0),
		scopes: make([]scope, 0),
	}
}

// Remove the last recorded token on tape.
// Should not be called during a rewind
func (t *tape) dropToken() {
	if t.record < 0 {
		return
	}
	top := len(t.tokens) - 1
	if top < 0 {
		return
	}
	t.tokens = t.tokens[:top]
}

// Return the next token which should be run
func (t *tape) nextToken() (token, bool, error) {
	top := len(t.scopes) - 1
	passive := false
	if top >= 0 {
		passive = t.scopes[top].lazy
	}

	for i := top; i >= 0; i-- {
		scope := t.scopes[i]
		if scope.rewinder < 0 {
			continue
		}
		next := scope.start + scope.rewinder
		tok := t.tokens[next]
		if next+1 < len(t.tokens) {
			t.scopes[i].rewinder += 1
		} else {
			t.scopes[i].rewinder = -1
		}

		if i == top {
			return tok, false, nil
		}
		return tok, passive, nil
	}

	tok, err := t.lexer.lex()
	if errors.Is(err, io.EOF) && top >= 0 {
		return token{}, false, ErrDanglingScope
	}
	if err != nil {
		return token{}, false, err
	}

	if t.record >= 0 {
		t.tokens = append(t.tokens, tok)
	}
	return tok, passive, nil
}

// Return the symbol of the top scope.
func (t *tape) Symbol() commands.Symbol {
	length := len(t.scopes)
	if length == 0 {
		return nil
	}
	return t.scopes[length-1].symbol
}

// Add a new scope to the stack.
// If needed is false, the scope may not record all the scope tokens
// If lazy is true, all commands will run as passive in scope
func (t *tape) PushScope(symbol commands.Symbol, needed, lazy bool) {
	start := len(t.tokens)
	for i := len(t.scopes) - 1; i >= 0; i-- {
		scope := t.scopes[i]
		if scope.rewinder < 0 {
			continue
		}
		start = scope.start + scope.rewinder
	}

	if t.record < 0 && needed {
		t.record = len(t.scopes)
	}
	t.scopes = append(t.scopes, scope{symbol, start, lazy, -1})
}

// Remove the top scope and return its symbol
func (t *tape) PopScope() commands.Symbol {
	top := len(t.scopes) - 1
	if top < 0 {
		return nil
	}
	symbol := t.scopes[top].symbol

	t.scopes = t.scopes[:top]
	if top <= t.record {
		t.record = -1
		t.tokens = t.tokens[:0] // do not release the memory for later uses
	}
	return symbol
}

// Start running commands from the start of the top scope
func (t *tape) Rewind() {
	top := len(t.scopes) - 1
	if top < 0 {
		return
	}
	if t.scopes[top].start >= len(t.tokens) {
		return
	}
	t.scopes[top].rewinder = 0
}
