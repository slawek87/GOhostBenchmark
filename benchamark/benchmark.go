/*
Package inspection implements a simple host inspection.

The inspection process sends GET requests to a given host and saves information like:
    * URL - url of given host
    * StatusCode - status code of response
    * Duration - time between sent request and received response

All data is saved to MongoDB.
 */

package benchamark

import (
	"net/http"
	"time"
	"sync"
	"github.com/slawek87/GOhostBenchmark/settings"
)

type BenchmarkData struct {
	URL         string
	StatusCode      int
	Duration    float64 // time in Seconds
}

type Benchmark struct {}


func (benchmark *Benchmark) sendRequest(url string) int {
	resp, _ := http.Get(url)
	return resp.StatusCode
}

func (benchmark *Benchmark) saveData(host *BenchmarkData) error {
	dbSession := settings.MongoDB()
	defer dbSession.Close()
	collection := dbSession.DB("StressTests").C("hosts")
	return collection.Insert(&host)
}

func (benchmark *Benchmark) measureDuration(url string) BenchmarkData {
	startTime := time.Now()
	statusCode := benchmark.sendRequest(url)
	duration := time.Since(startTime).Seconds()
	return BenchmarkData{url, statusCode, duration}
}

// use this method to run benchamark for current host.
// requestsNumber is a number of GET requests sent to given url.
func (benchmark *Benchmark) Process(url string, requestsNumber int) {
	requests := make(chan int)
	var waitGroup sync.WaitGroup

	waitGroup.Add(requestsNumber)

	for i := 1; i <= requestsNumber; i++ {
		go func() {
			defer waitGroup.Done()
			duration := benchmark.measureDuration(url)
			benchmark.saveData(&duration)
			requests <- i
		}()
	}

	waitGroup.Wait()
}

func (benchmark *Benchmark) Run(url string, requestsNumber int) {
	scan := Scanner{}
	url = normalizeUrl(url)

	// bench all internal urls
	for _, url := range scan.ScanHost(url) {
		benchmark.Process(url, requestsNumber)
	}
	// bench tha main url.
	benchmark.Process(url, requestsNumber)
}