package main

import (
	"encoding/xml"
	"log"
	"strings"
)

type feed struct {
	Doc []Doc `xml:"doc"`
}

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

	data := `
	<feed>
	<doc>
	<title>Wikipedia: Mariapia Degli Esposti</title>
	<url>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti</url>
	<abstract>Cancer Council of Western Australia Cancer Researcher of the Year 2017Research Excellence Awards - Cancer Council Western Australia (2017)</abstract>
	<links>
	<sublink linktype="nav"><anchor>Research</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#Research</link></sublink>
	<sublink linktype="nav"><anchor>Education</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#Education</link></sublink>
	</links>
	</doc>
	</feed>
	`
	mariapia := &feed{}
	xml.Unmarshal([]byte(data), mariapia)

	log.Printf("%#v\n", mariapia)
	log.Println("title:", mariapia.Doc[0].Title)
	log.Println("url:", mariapia.Doc[0].Url)
	log.Println("abstract:", mariapia.Doc[0].Abstract)
	log.Println("links:", len(strings.Split(mariapia.Doc[0].Links, "\n"))/2)

}
