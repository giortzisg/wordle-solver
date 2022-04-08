package config

import (
	_ "embed"
	"encoding/json"
	"log"

	"github.com/giortzisg/wordle-solver/internal/entities"
)

//go:embed words.json
var wordsFile []byte

//go:embed common_words.json
var commonWordsFile []byte

//go:embed probability_map.json
var ProbabilityMap []byte

//go:embed config.json
var configMode []byte

type configMap struct {
	HardMode bool `json:"hardMode"`
}

func LoadWords() entities.Words {
	var cfg configMap
	err := json.Unmarshal(configMode, &cfg)
	if err != nil {
		log.Fatalf("cannot unmarshal json object: %v", err)
	}
	var wFile []byte
	if cfg.HardMode {
		wFile = wordsFile
	} else {
		wFile = commonWordsFile
	}
	var words entities.Words
	err = json.Unmarshal(wFile, &words)
	if err != nil {
		log.Fatalf("cannot unmarshal json object: %v", err)
	}
	return words
}
