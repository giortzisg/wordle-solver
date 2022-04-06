package guess

import (
	"errors"
	"strings"

	"github.com/giortzisg/wordle-solver/internal/entities"
)

type letterHint struct {
	letter   rune
	included bool
	position bool
}

type letters []letterHint

// Guess contains the hints that wordle gives when entering a word.
type Guess struct {
	letters
}

func (g Guess) String() string {
	var s strings.Builder
	for _, l := range g.letters {
		s.WriteByte(byte(l.letter))
	}
	return s.String()
}

func (g Guess) StringPattern() string {
	var s strings.Builder
	for _, l := range g.letters {
		if l.position {
			s.WriteRune('g')
		} else if l.included {
			s.WriteRune('y')
		} else {
			s.WriteRune('w')
		}
	}
	return s.String()
}

func (g Guess) HintPattern() string {
	var s strings.Builder
	for _, c := range g.letters {
		if c.position {
			s.WriteRune('🟩')
		} else if c.included {
			s.WriteRune('🟨')
		} else {
			s.WriteString("⬜️")
		}
	}
	return s.String()
}

// Probability calculates the probability of a specific guess on a specific word set
func (g Guess) Probability(words entities.Words) float64 {
	possibleMatches := g.FilterWords(words)
	return float64(len(possibleMatches)) / float64(len(words))
}

func (g Guess) FilterWords(words entities.Words) (possibleMatches entities.Words) {
	for _, w := range words {
		shouldFilter := false
		for i, hint := range g.letters {
			if hint.position {
				if w[i] != byte(hint.letter) {
					shouldFilter = true
					break
				}
			} else if hint.included {
				if !strings.ContainsRune(w, hint.letter) || strings.Count(w, string(hint.letter)) < strings.Count(g.String()[:i+1], string(hint.letter)) || w[i] == byte(hint.letter) {
					shouldFilter = true
					break
				}
			} else if strings.ContainsRune(w, hint.letter) {
				shouldFilter = true
				break
			}
		}
		if !shouldFilter {
			possibleMatches = append(possibleMatches, w)
		}
	}
	return possibleMatches
}

func New(wordGuessed string, hint string) (Guess, error) {
	var res Guess
	if len(wordGuessed) != 5 {
		return res, errors.New("invalid word length")
	}
	if len(hint) != 5 {
		return res, errors.New("invalid hint length")
	}
	for i, g := range wordGuessed {
		res.letters = append(res.letters, letterHint{letter: g, included: hint[i] == 'y' || hint[i] == 'g', position: hint[i] == 'g'})
	}
	return res, nil
}
