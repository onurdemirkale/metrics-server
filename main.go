package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	port            = "12345"
	metricsPath     = "/metrics"
	metricsFilePath = "metrics_from_special_app.txt"
	cacheDuration   = 1 * time.Minute
)

var (
	metrics     []byte // simple cache implementation
	lastUpdated time.Time
)

func main() {
	httpServer := &http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc(metricsPath, metricsHandler)

	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}

func metricsHandler(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	// read metrics file if cache is expired
	if time.Since(lastUpdated) > cacheDuration {

		var err error
		metrics, err = ioutil.ReadFile(metricsFilePath)
		if err != nil {
			http.Error(responseWriter, "Error reading metrics file", http.StatusInternalServerError)
			return
		}

		lastUpdated = time.Now()
	}

	// write metrics to response writer
	responseWriter.Write(metrics)
}
