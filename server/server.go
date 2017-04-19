package server

import (
	"log"
	"net/http"
	"strconv"
)

// Port is the http port the server is started on.
var Port = 6060

// Start starts up the http rest server.
func Start() {
	http.HandleFunc(dogFosterURLPath, DogFosterHandler)
	startServer()
}

func startServer() {
	err := http.ListenAndServe(":"+strconv.Itoa(Port), nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
