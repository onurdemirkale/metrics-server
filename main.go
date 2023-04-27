package main

import (
	"log"
	"net/http"
)

func main() {
	httpServer := &http.Server{
		Addr: ":12345",
	}

	err := httpServer.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
