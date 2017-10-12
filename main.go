package main

import "github.com/slawek87/GOhostBenchmark/benchamark"

func main() {
    bench := benchamark.Benchmark{}
    bench.Run("http://145.239.91.118", 1000)
}