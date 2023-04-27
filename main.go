package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	port            = "12345"
	metricsPath     = "/metrics"
	metricsFilePath = "metrics_from_special_app.txt"
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

	// read metric values from file
	metricsData, err := ioutil.ReadFile(metricsFilePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parse metrics data into a map
	metrics := make(map[string]string)

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

	// write metrics to response writer
	for key, value := range metrics {
		fmt.Fprintf(w, "%s=%s\n", key, value)
	}
}
