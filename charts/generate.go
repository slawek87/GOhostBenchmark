package charts

import (
	"github.com/slawek87/GOhostBenchmark/settings"
	"github.com/slawek87/GOhostBenchmark/benchamark"
	"gopkg.in/mgo.v2/bson"
	goChart "github.com/wcharczuk/go-chart"
	"net/http"
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
	url := "http://145.239.91.118"
	data := chart.getDataForURL(url)
	x, y := chart.prepareXY(data)

	graph := goChart.Chart{
		Series: []goChart.Series{
			goChart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(goChart.PNG, res)
}
