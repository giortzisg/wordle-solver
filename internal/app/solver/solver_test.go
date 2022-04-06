package solver

import (
	"errors"
	"fmt"
	"testing"

	"github.com/giortzisg/wordle-solver/internal/entities"
	"github.com/stretchr/testify/assert"
)

func Test_Solve(t *testing.T) {
	tests := map[string]struct {
		words       entities.Words
		wordGuessed string
		hints       string
		expWords    entities.Words
		expErr      error
	}{
		"Correctly solves wordle": {
			words:       entities.Words{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "acaca"},
			wordGuessed: "acaca",
			hints:       "gwgwg",
			expWords:    entities.Words{"aaaaa"},
			expErr:      nil,
		},
		"Error on invalid word length": {
			words:       entities.Words{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "acaca"},
			wordGuessed: "acacaa",
			hints:       "gwgwg",
			expWords:    entities.Words{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "acaca"},
			expErr:      fmt.Errorf("cannot solve wordle: %w", errors.New("invalid word length")),
		},
		"Error on invalid hint length": {
			words:       entities.Words{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "acaca"},
			wordGuessed: "acaca",
			hints:       "gwgwgd",
			expWords:    entities.Words{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "acaca"},
			expErr:      fmt.Errorf("cannot solve wordle: %w", errors.New("invalid hint length")),
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			w, err := Solve(tt.wordGuessed, tt.hints, tt.words)
			assert.Equal(t, tt.expWords, w)
			assert.Equal(t, tt.expErr, err)
		})
	}
}
