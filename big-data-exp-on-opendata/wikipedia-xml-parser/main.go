package main

import (
	"log"
)

func main() {
	sourcefile := "./enwiki-20210620-abstract.xml"

	db := openDatabase()
	defer db.Close()

	file := openXML(sourcefile)
	defer file.Close()

	startDatabase(db)
	parseAndSave(file, db)

	log.Println("Parsing complete.")
}
