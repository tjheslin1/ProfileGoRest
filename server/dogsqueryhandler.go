package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tjheslin1/ProfileGoRest/dogrepo"
)

// DogFosterHandler provides a function for http.HandlerFunc
type DogFosterHandler struct {
	URLPath string
	logger  *log.Logger
}

// DogFosterHandler handles requests to query the dog repository
func (dfh *DogFosterHandler) DogFosterHandler(w http.ResponseWriter, req *http.Request) {
	dfh.logger.Printf("%v\n::::::\n", req)

	repository := dogrepo.DogFosterRepository{}

	dog := repository.ExampleDogs()

	data, _ := json.Marshal(&dog)

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
