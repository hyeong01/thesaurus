package thesaurus

import (
	"github.com/hyeong01/thesaurus/internal/api/models"
	"github.com/hyeong01/thesaurus/internal/db"
)

func GetSynonyms(word string) (models.Synonym, error) {
	synonyms, similarities, err := db.GetSynonymsAndSimilarities(word)
	if err != nil {
		return models.Synonym{}, err
	}

	return models.Synonym{
		Word:         word,
		Synonyms:     synonyms,
		Similarities: similarities,
	}, nil
}
