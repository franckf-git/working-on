package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	db := openDatabase()
	defer db.Close()
	startDatabase(db)

	var datas = getAlphabetFreq(db, 0)
	createChart(datas)
}

func getAlphabetFreq(db *sql.DB, limit int) map[string]float64 {
	if limit == 0 {
		limit = countValues(db)
	}

	alphabet := make(map[string]float64)

	rows, err := db.Query("SELECT title,links FROM doc LIMIT ?", limit)
	if err != nil {
		log.Fatal("Select fail - executing query:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var title string
		var links int
		err = rows.Scan(&title, &links)
		if err != nil {
			log.Fatal("Select fail - scanning values:", err)
		}
		letter := parsingTitle(title)
		alphabet[letter] = float64(links) + alphabet[letter]
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Select fail - reading rows:", err)
	}

	return alphabet
}

func createChart(datas map[string]float64) {

	values := []chart.Value{}

	for label, value := range datas {
		adding := []chart.Value{{Label: label, Value: value}}
		values = append(values, adding...)
	}

	graph := chart.BarChart{
		Title: "Stats - Links by Letter",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   1024,
		Width:    4096,
		BarWidth: 60,
		Bars:     values,
	}

	chartFile, _ := os.Create("output.png")
	defer chartFile.Close()
	graph.Render(chart.PNG, chartFile)
}

func parsingTitle(title string) string {
	withoutWikipedia := strings.Replace(title, "Wikipedia: ", "", -1)
	return withoutWikipedia[:1]
}
