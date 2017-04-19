package server

import (
	"encoding/json"
	"net/http"

	"github.com/tjheslin1/ProfileGoRest/dogrepo"
)

var dogsQueryURLPath = "/foster"

// DogFosterHandler handles requests to query the dog repository
func DogFosterHandler(w http.ResponseWriter, req *http.Request) {
	repository := dogrepo.DogFosterRepository{}

	dog := repository.ExampleDogs()

	data, _ := json.Marshal(&dog)

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
