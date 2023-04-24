package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hyeong01/thesaurus/internal/thesaurus"
)

func GetSynonyms(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("word")
	print(query)

	if query == "" {
		http.Error(w, "Please provide a word", http.StatusBadRequest)
		return
	}

	synonyms, err := thesaurus.GetSynonyms(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(synonyms)
}
