package models

type Synonym struct {
	Word         string   `json:"word"`
	Synonyms     []string `json:"synonyms"`
	Similarities []int    `json:"similarities"`
}
