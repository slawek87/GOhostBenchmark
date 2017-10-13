package main

import (
	//"github.com/slawek87/GOhostBenchmark/charts"
	//"net/http"
	//"github.com/slawek87/GOhostBenchmark/benchamark"
	"github.com/slawek87/GOhostBenchmark/charts"
	"net/http"
)

func main() {
	//bench := benchmark.Benchmark{}
	//bench.Run("http://147.135.208.25", 1000)
	chart := charts.Chart{}
	render := chart.Render

	http.HandleFunc("/", render)
	http.ListenAndServe(":8080", nil)
}
