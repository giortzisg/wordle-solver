package solver

import (
	"strings"
)

type letterHint struct {
	letter   rune
	included bool
	position bool
}

type wordHint struct {
	hints []letterHint
}

func (h wordHint) Word() string {
	var s strings.Builder
	for _, c := range h.hints {
		s.WriteByte(byte(c.letter))
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

func FilterWords(words []string, guess wordHint) (possibleWords []string) {
	for _, w := range words {
		shouldFilter := false
		for i, hint := range guess.hints {
			if hint.position {
				if w[i] != byte(hint.letter) {
					shouldFilter = true
					break
				}
			} else if hint.included {
				if !strings.ContainsRune(w, hint.letter) {
					shouldFilter = true
					break
				}
			} else if strings.ContainsRune(w, hint.letter) {
				shouldFilter = true
				break
			}
		}
		if !shouldFilter {
			possibleWords = append(possibleWords, w)
		}
	}
	return possibleWords
}

func NewGuess(guess string, hints string) (res wordHint) {
	for i, g := range guess {
		res.hints = append(res.hints, letterHint{letter: g, included: hints[i] == 'y', position: hints[i] == 'g'})
	}
	return res
}
