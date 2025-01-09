package main

import (
	"log"

	"github.com/iChemy/MyKnoQ/backend/infra/db"
)

func main() {
	_, err := db.Setup()

	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
}
