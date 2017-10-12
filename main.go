package main

import (
	"github.com/slawek87/GOhostBenchmark/charts"
	"net/http"
)

func main() {
	//bench := benchamark.Benchmark{}
	//bench.Run("http://145.239.91.118", 1000)
	chart := charts.Chart{}
	render := chart.Render

	http.HandleFunc("/", render)
	http.ListenAndServe(":8080", nil)
}
