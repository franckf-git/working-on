package main

import (
	"log"
)

func main() {
	db := openDatabase()
	defer db.Close()

	startDatabase(db)
	insertDatabase(db, 8, "title", "https://url.test", "abstract", 2)
	insertDatabase(db, 9, "dex", "url", "abstract", 54)
	insertDatabase(db, 11, "TITLE", "https://url.fr", "abstract4", 462)
	output := selectAllDatabase(db)
	log.Println("Results:", output)
}
