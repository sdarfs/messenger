package main

import (
	"log"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal("could not connect to db")
	}
}
