package main

import (
	"log"
	"net/http"

	"github.com/hyeong01/thesaurus/config"
	"github.com/hyeong01/thesaurus/internal/api/router"
	"github.com/hyeong01/thesaurus/internal/db"
)

func main() {
	config.LoadConfig()
	db.Connect()

	r := router.NewRouter()
	port := config.GetConfig().Port

	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
