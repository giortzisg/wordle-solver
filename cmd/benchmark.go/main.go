package main

import (
	"fmt"
	"log"
	"strings"

	_ "embed"

	"github.com/giortzisg/wordle-solver/config"
	"github.com/giortzisg/wordle-solver/internal/app/solver"
	"github.com/schollz/progressbar/v3"
)

// evaluateGuess returns a hint based on the given solution
func evaluateGuess(guess string, solution string) (hint string) {
	for i, letter := range guess {
		lh := "w"
		if strings.ContainsRune(solution, letter) && strings.Count(solution, string(letter)) >= strings.Count(guess[:i+1], string(letter)) {
			lh = "y"
		}
		if letter == rune(solution[i]) {
			lh = "g"
		}
		hint = hint + lh
	}
	return hint
}

func main() {
	var solutions, count float64
	words := config.LoadWords()
	bar := progressbar.Default(int64(len(words)))
	for _, w := range words {
		err := bar.Add(1)
		if err != nil {
			log.Fatalf("progress bar error: %v", err)
		}
		wordGuess := "raise"
		tries := 0.0
		testWords := words
		count++
		for {
			tries++
			hint := evaluateGuess(wordGuess, w)
			testWords, err = solver.Solve(wordGuess, hint, testWords)
			if err != nil {
				log.Fatal(err)
			}
			if len(testWords) == 1 {
				solutions += tries
				break
			} else {
				wordGuess = testWords[0]
			}
		}
	}
	fmt.Printf("The accuracy of wordle solver is %v\n", solutions/count)
}
