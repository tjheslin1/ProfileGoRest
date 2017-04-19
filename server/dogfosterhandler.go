package server

import (
	"encoding/json"
	"net/http"

	"github.com/tjheslin1/ProfileGoRest/dogrepo"
)

var dogFosterURLPath = "/foster"

// DogFosterHandler handles requests to query the dog repository
func DogFosterHandler(w http.ResponseWriter, req *http.Request) {
	repository := dogrepo.DogFosterRepository{}

	dog := repository.ExampleDog()

	data, _ := json.Marshal(&dog)

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
