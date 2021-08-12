package main

import (
	"bufio"
	"encoding/xml"
	"log"
	"os"
	"strings"
)

type Doc struct {
	Title    string `xml:"title"`
	Url      string `xml:"url"`
	Abstract string `xml:"abstract"`
	Links    string `xml:"links"`
}

func main() {
	db := openDatabase()
	defer db.Close()
	startDatabase(db)

	file, err := os.Open("./small.xml")
	if err != nil {
		log.Println("Error opening xml file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var block string
	incrementId := 1
	for scanner.Scan() {
		if scanner.Text() == "<doc>" {
			block = ""
		}

		block = block + scanner.Text()

		if scanner.Text() == "</doc>" {
			blockParsing := &Doc{}
			xml.Unmarshal([]byte(block), blockParsing)
			countLinks := len(strings.Split(blockParsing.Links, "\n")) / 2

			insertDatabase(db, incrementId, blockParsing.Title, blockParsing.Url, blockParsing.Abstract, countLinks)

			incrementId = incrementId + 1
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading scan of file:", err)
	}

	log.Println("Parsing complete.")
	log.Println(selectAllDatabase(db))
}
