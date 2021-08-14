package main

import (
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

func main() {
	createChart()
}

func createChart() {

	values := []chart.Value{}

	adding1 := []chart.Value{{Value: 50, Label: "Encore"}}
	adding2 := []chart.Value{{Value: 10, Label: "Avant"}}
	values = append(values, adding1...)
	values = append(values, adding2...)

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
