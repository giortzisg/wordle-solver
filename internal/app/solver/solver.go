package solver

import (
	_ "embed"
	"fmt"
	"math"
	"sort"

	"github.com/giortzisg/wordle-solver/internal/app/guess"
	"github.com/giortzisg/wordle-solver/internal/entities"
)

func calculateWordEntropy(words entities.Words, wordGuessed string) float64 {
	possibleHints := []string{"g", "y", "w"}
	entropy := 0.0
	for _, r1 := range possibleHints {
		for _, r2 := range possibleHints {
			for _, r3 := range possibleHints {
				for _, r4 := range possibleHints {
					for _, r5 := range possibleHints {
						hint := fmt.Sprintf("%v%v%v%v%v", r1, r2, r3, r4, r5)
						g, err := guess.New(wordGuessed, hint)
						if err != nil {
							return 0
						}
						p := g.Probability(words)
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

func sortWords(words entities.Words) entities.Words {
	wordsWithEntropy := make(map[string]float64)
	for _, w := range words {
		wordsWithEntropy[w] = calculateWordEntropy(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		return wordsWithEntropy[words[i]] > wordsWithEntropy[words[j]]
	})
	return words
}

func Solve(wordGuessed string, hints string, words entities.Words) (entities.Words, error) {
	g, err := guess.New(wordGuessed, hints)
	if err != nil {
		return words, fmt.Errorf("cannot solve wordle: %w", err)
	}
	words = g.FilterWords(words)
	return sortWords(words), nil
}
