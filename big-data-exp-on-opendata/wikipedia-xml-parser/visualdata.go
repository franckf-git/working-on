package main

import (
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	var datas = map[string]float64{
		"A": 25,
		"B": 36,
		"C": 0,
		"D": 8,
	}
	createChart(datas)
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
		Height:   512,
		BarWidth: 60,
		Bars:     values,
	}

	chartFile, _ := os.Create("output.png")
	defer chartFile.Close()
	graph.Render(chart.PNG, chartFile)
}
