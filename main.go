package main

import (
	"bytes"
	"fmt"
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
	metrics     map[string]string // simple cache implementation
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

func metricsHandler(w http.ResponseWriter, r *http.Request) {

	// read metrics file if cache is expired
	if time.Since(lastUpdated) > cacheDuration {

		// read metric values from file
		metricsData, err := ioutil.ReadFile(metricsFilePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// parse metrics data into a map
		metrics = make(map[string]string)

		for _, line := range bytes.Split(metricsData, []byte("\n")) {
			if len(line) == 0 {
				continue
			}
			parts := bytes.SplitN(line, []byte("="), 2)
			if len(parts) != 2 {
				continue
			}
			metrics[string(parts[0])] = string(parts[1])
		}

		lastUpdated = time.Now()
	}

	// write metrics to response writer
	for key, value := range metrics {
		fmt.Fprintf(w, "%s=%s\n", key, value)
	}
}
