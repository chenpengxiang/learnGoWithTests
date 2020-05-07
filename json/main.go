package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, _ := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatalf("could not listen no port 8000 %v", err)
	}
}