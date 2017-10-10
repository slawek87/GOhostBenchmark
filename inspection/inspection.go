/*
Package inspection implements a simple host inspection.

The inspection process sends GET requests to a given host and saves information like:
    * URL - url of given host
    * StatusCode - status code of response
    * Duration - time between sent request and received response

All data is saved to MongoDB.
 */

package inspection

import (
	"net/http"
	"time"
	"sync"
	"github.com/slawek87/GOhostBenchmark/settings"
)

type Host struct {
	URL         string
	StatusCode      int
	Duration    float64 // time in Seconds
}

type InspectHost struct {}


func (inspection *InspectHost) sendRequest(url string) int {
	resp, _ := http.Get(url)
	return resp.StatusCode
}

func (inspection *InspectHost) saveData(host *Host) error {
	dbSession := settings.MongoDB()
	defer dbSession.Close()
	collection := dbSession.DB("StressTests").C("hosts")
	return collection.Insert(&host)
}

func (inspection *InspectHost) measureDuration(url string) Host {
	startTime := time.Now()
	statusCode := inspection.sendRequest(url)
	duration := time.Since(startTime).Seconds()
	return Host{url, statusCode, duration}
}

// use this method to run inspection for current host.
// requestsNumber is a number of GET requests sent to given url.
func (inspection *InspectHost) Inspect(url string, requestsNumber int) {
	requests := make(chan int)
	var waitGroup sync.WaitGroup

	waitGroup.Add(requestsNumber)

	for i := 1; i <= requestsNumber; i++ {
		go func() {
			defer waitGroup.Done()
			duration := inspection.measureDuration(url)
			inspection.saveData(&duration)
			requests <- i
		}()
	}

	waitGroup.Wait()
}
