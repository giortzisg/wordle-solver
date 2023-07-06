package solver

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"sort"
	"sync"

	"github.com/giortzisg/wordle-solver/config"
	"github.com/giortzisg/wordle-solver/internal/app/guess"
	"github.com/giortzisg/wordle-solver/internal/entities"
)

var probabilityMap map[string]float64
var possibleHints []string

func init() {
	probabilityFile := config.ProbabilityMap
	possibleHints = createAllPossibleHints()
	err := json.Unmarshal(probabilityFile, &probabilityMap)
	if err != nil {
		log.Fatalf("cannot unmarshal json object: %v", err)
	}
}

func createAllPossibleHints() []string {
	res := []string{}
	possibleHints := []string{"g", "y", "w"}
	for _, r1 := range possibleHints {
		for _, r2 := range possibleHints {
			for _, r3 := range possibleHints {
				for _, r4 := range possibleHints {
					for _, r5 := range possibleHints {
						hint := fmt.Sprintf("%v%v%v%v%v", r1, r2, r3, r4, r5)
						res = append(res, hint)
					}
				}
			}
		}
	}
	return res
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func calculateWordEntropy(words entities.Words, wordGuessed string) float64 {
	entropy := 0.0
	p := 0.0
	for _, hint := range possibleHints {
		g, _ := guess.New(wordGuessed, hint)
		p += g.Probability(words) * probabilityMap[wordGuessed]
	}

	if p != 0 {
		entropy += p * math.Log2(1/p)
	}

	return entropy
}

func sortWords(words entities.Words) entities.Words {
	var wg sync.WaitGroup
	var wordsChunks [][]string

	wordsWithEntropy := make(map[string]float64, len(words))

	if len(words) < 4 {
		wordsChunks = chunkSlice(words, 1)
	} else {
		wordsChunks = chunkSlice(words, len(words)/4)
	}

	// spawn go routines to handle the words quicker
	wordsCh := make(chan struct {
		word        string
		probability float64
	}, len(words))

	for i := 0; i < len(wordsChunks); i++ {
		wg.Add(1)
		chunk := wordsChunks[i]
		go func() {
			defer wg.Done()
			for _, w := range chunk {
				wordsCh <- struct {
					word        string
					probability float64
				}{w, calculateWordEntropy(words, w)}
			}
		}()
	}
	wg.Wait()
	close(wordsCh)

	for ch := range wordsCh {
		wordsWithEntropy[ch.word] = ch.probability
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
	// filter the words
	words = g.FilterWords(words)
	// sort the words based on their probability
	return sortWords(words), nil
}
