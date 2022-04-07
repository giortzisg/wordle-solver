package entities

import _ "embed"

//go:embed words.json
var WordsFile []byte

// Words struct which contains an array of words
type Words []string
