package main

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/giortzisg/wordle-solver/internal/app/solver"
	"github.com/giortzisg/wordle-solver/internal/entities"
)

//go:embed internal/entities/words.json
var wordsFile []byte

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input entities.Words
	err := json.Unmarshal(wordsFile, &input)
	if err != nil {
		log.Fatalf("cannot unmarshal json object: %v", err)
	}
	fmt.Println(input.Words[0:5])
	for {
		fmt.Printf("Enter the word that you guessed\n")
		wordGuessed, _ := reader.ReadString('\n')
		fmt.Printf("Total words: %v\n", len(input.Words))
		fmt.Printf("Enter the response to the clue:\n - \033[1mW\033[0m: White \n - \033[33mY\033[0m: Yellow \n - \033[32mG\033[0m: Green \n")
		hints, _ := reader.ReadString('\n')
		guess := solver.NewGuess(wordGuessed, hints)
		input = solver.FilterWords(input, guess)
		sorted := solver.SortPossibleWords(input)
		if len(sorted.Words) == 1 {
			fmt.Printf("The solution is: %s\n", sorted.Words[0])
			break
		} else {
			if len(sorted.Words) < 5 {
				fmt.Printf("There are currently %d possible words. Some of the best moves are to try \033[1m%s\033[0m\n", len(sorted.Words), sorted.Words)
			} else {
				fmt.Printf("There are currently %d possible words. Some of the best moves are to try \033[1m%s\033[0m\n", len(sorted.Words), sorted.Words[0:5])
			}
		}

	}
}
