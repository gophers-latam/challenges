package main

import (
	"github.com/tomiok/challenge-svc/storage"
	"os"
)

func main() {
	storage.Migrate()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	start(port)
}
