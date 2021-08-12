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
	block := ""
	incrementId := 1

	for scanner.Scan() {
		line := scanner.Text()

		if line == "<doc>" {
			block = ""
		}
		if line == "</doc>" {
			title, url, abstract, links := parseBlock(block)
			insertDatabase(db, incrementId, title, url, abstract, links)
			incrementId = incrementId + 1
		}
		block = block + line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading scan of file:", err)
	}

	log.Println("Parsing complete.")
	log.Println(selectAllDatabase(db))
}

func parseBlock(block string) (title string, url string, abstract string, countLinks int) {
	blockParsing := &Doc{}
	xml.Unmarshal([]byte(block), blockParsing)

	title = blockParsing.Title
	url = blockParsing.Url
	abstract = blockParsing.Abstract
	countLinks = len(strings.Split(blockParsing.Links, "\n")) / 2
	return
}
