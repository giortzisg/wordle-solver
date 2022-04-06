package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/giortzisg/wordle-solver/internal/app/solver"
)

//go:embed internal/entities/words.txt
var wordsFile []byte

func main() {
	reader := bufio.NewReader(os.Stdin)
	words := make([]string, 0, 6000)
	// Clean
	{
		dirtyWords := strings.Split(string(wordsFile), "\n")
		for _, w := range dirtyWords {
			if len(w) != 5 {
				continue
			}
			words = append(words, w)
		}
	}
	for {
		fmt.Printf("Enter the word that you guessed\n")
		wordGuessed, _ := reader.ReadString('\n')
		fmt.Printf("Enter the response to the clue:\n - \033[1mW\033[0m: White \n - \033[33mY\033[0m: Yellow \n - \033[32mG\033[0m: Green \n")
		hints, _ := reader.ReadString('\n')
		guess := solver.NewGuess(wordGuessed, hints)
		words = solver.FilterWords(words, guess)
		if len(words) == 1 {
			fmt.Printf("The solution is: %s\n", words[0])
			break
		} else {
			if len(words) < 5 {
				fmt.Printf("There are currently %d possible words. Some of the best moves are to try \033[1m%s\033[0m\n", len(words), words)
			}
			fmt.Printf("There are currently %d possible words. Some of the best moves are to try \033[1m%s\033[0m\n", len(words), words[0:5])
		}

	}
}
