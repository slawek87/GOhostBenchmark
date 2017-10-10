package main

import "github.com/slawek87/GOhostBenchmark/inspection"

func main() {
    insp := inspection.InspectHost{}
    insp.Inspect("http://147.135.208.25/watch/20", 1000)
}