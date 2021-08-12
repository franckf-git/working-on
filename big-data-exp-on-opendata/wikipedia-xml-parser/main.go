package main

import (
	"log"
)

func main() {
	sourcefile := "./small.xml"

	db := openDatabase()
	defer db.Close()

	file := openXML(sourcefile)
	defer file.Close()

	startDatabase(db)
	parseAndSave(file, db)

	log.Println("Parsing complete.")
	log.Println(selectAllDatabase(db))
}
