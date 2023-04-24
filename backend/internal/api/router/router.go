package router

import (
	"net/http"

	"github.com/hyeong01/thesaurus/internal/api/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/synonyms", handlers.GetSynonyms)

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/", fs)

	return handlers.CORS(mux)
}
