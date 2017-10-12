package charts

import (
	"github.com/slawek87/GOhostBenchmark/settings"
	"github.com/slawek87/GOhostBenchmark/benchamark"
	"gopkg.in/mgo.v2/bson"
	goChart "github.com/wcharczuk/go-chart"
	"net/http"
	//"fmt"
)

type Chart struct {}


// returns data for given url.
func (chart *Chart) getDataForURL(url string) []benchamark.BenchmarkData {
	var result []benchamark.BenchmarkData

	url = settings.NormalizeUrl(url)

	dbSession := settings.MongoDB()
	defer dbSession.Close()

	collection := dbSession.DB("StressTests").C("hosts")
	collection.Find(bson.M{"url": url}).All(&result)

	return result
}

func (chart *Chart) prepareXY(data []benchamark.BenchmarkData) ([]float64, []float64) {
	var x, y []float64

	for key, value := range data {
		x = append(x, float64(key))
		y = append(y, value.Duration)
	}

	return x, y
}

func (chart *Chart) Render(res http.ResponseWriter, req *http.Request) {
	url := "http://147.135.208.25"
	data := chart.getDataForURL(url)
	x, y := chart.prepareXY(data)

	graph := goChart.Chart{
		XAxis: goChart.XAxis{
			Name:      "Request Number",
			NameStyle: goChart.StyleShow(),
			Style:     goChart.StyleShow(),
		},
		YAxis: goChart.YAxis{
			Name:      "Response Time",
			NameStyle: goChart.StyleShow(),
			Style:     goChart.StyleShow(),
		},
		Series: []goChart.Series{
			goChart.ContinuousSeries{
				Style: goChart.Style{
					Show:        true,
					StrokeColor: goChart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   goChart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: x,
				YValues: y,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(goChart.PNG, res)
}
