package main

import "os"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	start(port)
}
