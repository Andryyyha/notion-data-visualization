package renderer

import (
	"math"
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

func generateBarData(data map[string]float64) ([]chart.Value, float64) {
	result := make([]chart.Value, 0)
	max := 0.0
	for k, v := range data {
		if data[k] != 0 {
			result = append(result, chart.Value{Value: v, Label: k})
			if data[k] > max {
				max = data[k]
			}
		}
	}
	return result, max
}

func Render(data map[string]float64) {
	barData, max := generateBarData(data)
	barChart := chart.BarChart{
		Bars: barData,
		YAxis: chart.YAxis{
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: math.Ceil(math.Floor(max)/100000) * 100000,
			},
		},
		// Canvas: chart.Style{
		// 	FillColor: drawing.Color{R: 200, G: 200, B: 200, A: 255},
		// },
	}
	f, _ := os.Create(os.Getenv("HOME") + "/notion-data-visualization.old/chart.svg")
	defer f.Close()
	barChart.Render(chart.SVG, f)
}
