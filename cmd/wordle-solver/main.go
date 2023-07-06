package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/giortzisg/wordle-solver/config"
	"github.com/giortzisg/wordle-solver/internal/app/solver"
)

func main() {
	corpus := config.LoadWords()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter the word that you guessed\n")
		wordGuessed, _ := reader.ReadString('\n')
		wordGuessed = strings.TrimSuffix(wordGuessed, "\n")
		// trim carriage return for Windows
		wordGuessed = strings.TrimSuffix(wordGuessed, "\r")
		fmt.Printf("Enter the response to the clue:\n - \033[1mW\033[0m: White \n - \033[33mY\033[0m: Yellow \n - \033[32mG\033[0m: Green \n")
		hint, _ := reader.ReadString('\n')
		hint = strings.TrimSuffix(hint, "\n")
		// trim carriage return for Windows
		hint = strings.TrimSuffix(hint, "\r")

		guesses, err := solver.Solve(wordGuessed, hint, corpus)
		if err != nil {
			log.Fatal(err)
		}
		if len(guesses) == 1 {
			fmt.Printf("The solution is: %s\n", guesses[0])
			break
		} else {
			if len(guesses) < 10 {
				fmt.Printf("There are currently %d possible words. Some of the best moves are to try \033[1m%s\033[0m\n", len(guesses), guesses)
			} else {
				fmt.Printf("There are currently %d possible words. Some of the best moves are to try \033[1m%s\033[0m\n", len(guesses), guesses[0:10])
			}
		}
		corpus = guesses
	}

}
