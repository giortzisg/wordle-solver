package solver

import (
	"testing"

	"github.com/giortzisg/wordle-solver/internal/entities"
	"github.com/stretchr/testify/assert"
)

func Test_EvaluateGuess(t *testing.T) {
	tests := map[string]struct {
		guess    string
		solution string
		expResp  wordHint
	}{
		"Correctly matches word": {
			guess:    "aabbb",
			solution: "aabbb",
			expResp: wordHint{
				hints: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: true, position: true},
					{letter: 'b', included: true, position: true},
					{letter: 'b', included: true, position: true},
				},
			},
		},
		"Marks correct number of existing letters not in correct spot": {
			guess:    "cdaaa",
			solution: "aabbb",
			expResp: wordHint{
				hints: []letterHint{
					{letter: 'c', included: false, position: false},
					{letter: 'd', included: false, position: false},
					{letter: 'a', included: true, position: false},
					{letter: 'a', included: true, position: false},
					{letter: 'a', included: false, position: false},
				},
			},
		},
		"Marks non existing letters": {
			guess:    "aabbc",
			solution: "aabbb",
			expResp: wordHint{
				hints: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: true, position: true},
					{letter: 'b', included: true, position: true},
					{letter: 'c', included: false, position: false},
				},
			},
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			resp := EvaluateGuess(tt.guess, tt.solution)
			assert.Equal(t, tt.expResp, resp)
		})
	}
}

func Test_FilterWords(t *testing.T) {
	tests := map[string]struct {
		words     entities.Words
		guessResp wordHint
		expWords  []string
	}{
		"Correctly returns words without included characters": {
			words: entities.Words{
				Words: []string{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "adada"},
			},
			guessResp: wordHint{
				hints: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'c', included: false, position: false},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: false, position: false},
					{letter: 'a', included: true, position: true},
				},
			},
			expWords: []string{"aaaaa", "adada"},
		},
		"Correctly returns words with included characters": {
			words: entities.Words{
				Words: []string{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc", "acaaa", "aaaca", "abaaa"},
			},
			guessResp: wordHint{
				hints: []letterHint{
					{letter: 'a', included: true, position: true},
					{letter: 'c', included: true, position: false},
					{letter: 'a', included: true, position: true},
					{letter: 'b', included: false, position: false},
					{letter: 'a', included: true, position: true},
				},
			},
			expWords: []string{"aaaca"},
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			resp := FilterWords(tt.words, tt.guessResp)
			assert.Equal(t, tt.expWords, resp.Words)
		})
	}
}

func Test_wordHintProbability(t *testing.T) {
	tests := map[string]struct {
		words     entities.Words
		guessResp wordHint
		expProb   float64
	}{
		"Correctly calculates probability": {
			words: entities.Words{
				Words: []string{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc"},
			},
			guessResp: wordHint{
				hints: []letterHint{
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
			p := wordHintProbability(tt.words, tt.guessResp)
			assert.Equal(t, tt.expProb, p)
		})
	}
}

func Test_CalculateWordEntropy(t *testing.T) {
	tests := map[string]struct {
		words      entities.Words
		guess      string
		expEntropy float64
	}{
		"Correctly calculates probability": {
			words: entities.Words{
				Words: []string{"aaaaa", "bbbbbb", "baaaa", "aaaab", "ccccc"},
			},
			guess:      "acaca",
			expEntropy: 2.321928094887362,
		},
	}

	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			p := CalculateWordEntropy(tt.words, tt.guess)
			assert.Equal(t, tt.expEntropy, p)
		})
	}
}
