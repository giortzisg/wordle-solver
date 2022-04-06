package solver

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/giortzisg/wordle-solver/internal/entities"
)

type letterHint struct {
	letter   rune
	included bool
	position bool
}

// wordHint contains the hints that wordle gives when entering a word.
type wordHint struct {
	hints []letterHint
}

func (h wordHint) String() string {
	var s strings.Builder
	for _, c := range h.hints {
		s.WriteByte(byte(c.letter))
	}
	return s.String()
}

func (h wordHint) Pattern() string {
	var s strings.Builder
	for _, c := range h.hints {
		if c.position {
			s.WriteRune('g')
		} else if c.included {
			s.WriteRune('y')
		} else {
			s.WriteRune('w')
		}
	}
	return s.String()
}

func (h wordHint) Hints() string {
	var s strings.Builder
	for _, c := range h.hints {
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

func EvaluateGuess(guess string, solution string) (r wordHint) {
	for i, letter := range guess {
		r.hints = append(r.hints, letterHint{
			letter,
			strings.ContainsRune(solution, letter) && strings.Count(solution, string(letter)) >= strings.Count(guess[:i+1], string(letter)),
			letter == rune(solution[i]),
		})
	}
	return r
}

func FilterWords(words entities.Words, guess wordHint) (possibleWords entities.Words) {
	for _, w := range words.Words {
		shouldFilter := false
		for i, hint := range guess.hints {
			if hint.position {
				if w[i] != byte(hint.letter) {
					shouldFilter = true
					break
				}
			} else if hint.included {
				if !strings.ContainsRune(w, hint.letter) || strings.Count(w, string(hint.letter)) < strings.Count(guess.String()[:i+1], string(hint.letter)) || w[i] == byte(hint.letter) {
					shouldFilter = true
					break
				}
			} else if strings.ContainsRune(w, hint.letter) {
				shouldFilter = true
				break
			}
		}
		if !shouldFilter {
			possibleWords.Words = append(possibleWords.Words, w)
		}
	}
	return possibleWords
}

func NewGuess(guess string, hints string) (res wordHint) {
	for i, g := range guess {
		if i > 4 {
			return res
		}
		res.hints = append(res.hints, letterHint{letter: g, included: hints[i] == 'y', position: hints[i] == 'g'})
	}
	return res
}

func CalculateWordEntropy(words entities.Words, guess string) float64 {
	possibleHints := []string{"g", "y", "w"}
	entropy := 0.0
	for _, r1 := range possibleHints {
		for _, r2 := range possibleHints {
			for _, r3 := range possibleHints {
				for _, r4 := range possibleHints {
					for _, r5 := range possibleHints {
						hint := fmt.Sprintf("%v%v%v%v%v", r1, r2, r3, r4, r5)
						p := wordHintProbability(words, NewGuess(guess, hint))
						if p != 0 {
							entropy += p * math.Log2(1/p)
						}
					}
				}
			}
		}
	}
	return entropy
}

func wordHintProbability(words entities.Words, guess wordHint) float64 {
	possibleMatches := FilterWords(words, guess)
	return float64(len(possibleMatches.Words)) / float64(len(words.Words))
}

func SortPossibleWords(words entities.Words) entities.Words {
	wordsWithEntropy := make(map[string]float64)
	for _, w := range words.Words {
		wordsWithEntropy[w] = CalculateWordEntropy(words, w)
	}
	sort.Slice(words.Words, func(i, j int) bool {
		return wordsWithEntropy[words.Words[i]] > wordsWithEntropy[words.Words[j]]
	})
	return words
}
