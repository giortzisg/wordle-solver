package guess

import (
	"errors"
	"testing"

	"github.com/giortzisg/wordle-solver/internal/entities"
	"github.com/stretchr/testify/assert"
)

func Test_FilterWords(t *testing.T) {
	tests := map[string]struct {
		words     entities.Words
		guessResp Guess
		expWords  entities.Words
	}{
		"Correctly returns words without included characters": {
			words: entities.Words{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "adada"},
			guessResp: Guess{
				letters: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'c', included: false, position: false},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: false, position: false},
					{letter: 'a', included: true, position: true},
				},
			},
			expWords: entities.Words{"aaaaa", "adada"},
		},
		"Correctly returns words with included and non included on the same letter": {
			words: entities.Words{"arear", "alert"},
			guessResp: Guess{
				letters: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'r', included: true, position: false},
					{letter: 'e', included: true, position: true},
					{letter: 'a', included: false, position: false},
					{letter: 'r', included: false, position: false},
				},
			},
			expWords: entities.Words{"alert"},
		},
		"Correctly returns words with included characters": {
			words: entities.Words{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "acaaa", "aaaca", "abaaa"},
			guessResp: Guess{
				letters: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'c', included: true, position: false},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: false, position: false},
					{letter: 'a', included: true, position: true},
				},
			},
			expWords: entities.Words{"aaaca"},
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			resp := tt.guessResp.FilterWords(tt.words)
			assert.Equal(t, tt.expWords, resp)
		})
	}
}

func Test_Probability(t *testing.T) {
	tests := map[string]struct {
		words     entities.Words
		guessResp Guess
		expProb   float64
	}{
		"Correctly calculates probability": {
			words: entities.Words{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc"},
			guessResp: Guess{
				letters: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'c', included: false, position: false},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: false, position: false},
					{letter: 'a', included: true, position: true},
				},
			},
			expProb: 0.2,
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			p := tt.guessResp.Probability(tt.words)
			assert.Equal(t, tt.expProb, p)
		})
	}
}

func Test_StringPattern(t *testing.T) {
	tests := map[string]struct {
		guessResp     Guess
		expStrPattern string
	}{
		"Correctly returns pattern": {
			guessResp: Guess{
				letters: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'c', included: false, position: false},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: false, position: false},
					{letter: 'a', included: true, position: false},
				},
			},
			expStrPattern: "gwgwy",
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			p := tt.guessResp.StringPattern()
			assert.Equal(t, tt.expStrPattern, p)
		})
	}
}

func Test_Pattern(t *testing.T) {
	tests := map[string]struct {
		guessResp     Guess
		expStrPattern string
	}{
		"Correctly returns pattern": {
			guessResp: Guess{
				letters: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'c', included: false, position: false},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: false, position: false},
					{letter: 'a', included: true, position: false},
				},
			},
			expStrPattern: "🟩⬜️🟩⬜️🟨",
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			p := tt.guessResp.HintPattern()
			assert.Equal(t, tt.expStrPattern, p)
		})
	}
}

func Test_New(t *testing.T) {
	tests := map[string]struct {
		wordGuessed string
		hint        string
		expGuess    Guess
		expErr      error
	}{
		"Correctly returns guess": {
			wordGuessed: "acaba",
			hint:        "gwgwy",
			expGuess: Guess{
				letters: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'c', included: false, position: false},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: false, position: false},
					{letter: 'a', included: true, position: false},
				},
			},
			expErr: nil,
		},
		"Error on incorrect word length": {
			wordGuessed: "acabaa",
			hint:        "gwgwy",
			expGuess:    Guess{},
			expErr:      errors.New("invalid word length"),
		},
		"Error on incorrect hint length": {
			wordGuessed: "acaba",
			hint:        "gwgwyy",
			expGuess:    Guess{},
			expErr:      errors.New("invalid hint length"),
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			g, err := New(tt.wordGuessed, tt.hint)
			assert.Equal(t, tt.expGuess, g)
			assert.Equal(t, tt.expErr, err)
		})
	}
}
