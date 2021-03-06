// Package server defins a REST server, exposed on the provided `Port`.
package server

import (
	"log"
	"net/http"
	"strconv"
)

// Port is the http port the server is started on.
var Port = 6060

// Start starts up the http rest server.
func Start(logger *log.Logger) {
	dfh := DogFosterHandler{"/foster", logger}
	http.HandleFunc(dfh.URLPath, dfh.DogFosterHandler)
	http.HandleFunc("/ready", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(204)
	})
	startServer(logger)
}

func startServer(logger *log.Logger) {
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(Port), nil)
		if err != nil {
			logger.Println(err)
			panic(err)
		}
	}()

	logger.Printf("Server started on port: %v\n", strconv.Itoa(Port))
}
