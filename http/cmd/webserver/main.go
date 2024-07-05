package main

import (
	game "gocker/http"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := game.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	defer close()

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	log.Printf("Server is starting on port %s\n", ":5000")

	newGame := game.NewTexasHoldem(game.BlindAlerterFunc(game.Alerter), store)

	server, _ := game.NewPlayerServer(store, newGame)

	if err := http.ListenAndServe(":5000", game.LoggingMiddleware(server)); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}