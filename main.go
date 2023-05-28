package main

import (
	"log"

	"github.com/luxarasis4/twittor/database"
	"github.com/luxarasis4/twittor/handlers"
)

func main() {
	if !database.CheckConnection() {
		log.Fatal("Without connection connection")

		return
	}

	handlers.Managers()
}
