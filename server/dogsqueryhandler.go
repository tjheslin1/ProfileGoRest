package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tjheslin1/ProfileGoRest/dogrepo"
)

// DogFosterHandler provides a http Handle to handle requests to path specified
// in the server configuration.
type DogFosterHandler struct {
	URLPath string
	logger  *log.Logger
}

// DogFosterHandler handles requests to query the dog repository
// by providing an implementation of the handler func to be passed
// into `http.HandlerFunc`
func (dfh *DogFosterHandler) DogFosterHandler(w http.ResponseWriter, req *http.Request) {
	dfh.logger.Printf("REQUEST: \n%v\n::::::\n", req)

	repository := dogrepo.DogFosterRepository{}

	dogs := repository.ExampleDogs()

	data, _ := json.Marshal(&dogs)

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)

	dfh.logger.Printf("RESPONSE: \n200\n%v\n%v\n::::::\n", w.Header(), dogs)
}
