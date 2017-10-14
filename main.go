package main

import (
	"github.com/slawek87/GOhostBenchmark/charts"
	"net/http"
	//"github.com/slawek87/GOhostBenchmark/benchmark"
	//"github.com/slawek87/GOhostBenchmark/charts"
	//"net/http"
)

func main() {
	//bench := benchmark.Benchmark{}
	//bench.Run("http://145.239.91.118", 1000)
	chart := charts.Chart{}
	render := chart.Render

	http.HandleFunc("/stats", render)
	//http.HandleFunc("/benchmark", render)
	http.ListenAndServe(":8080", nil)
}
