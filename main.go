package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	entry "go1x_chi.template/entry"
)

func main() {
	godotenv.Load()

	r := entry.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, r)
}
