package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/tjheslin1/ProfileGoRest/dogrepo"
)

func TestMain(m *testing.M) {
	mockServer := setUpMockServer()
	defer mockServer.Close()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestDogFosterHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:6060/foster", nil)
	w := httptest.NewRecorder()

	dfh := DogFosterHandler{"/foster", log.New(os.Stdout, "", 0)}
	dfh.DogFosterHandler(w, req)

	if err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Error(err)
	}

	w.Body.Read(data)

	var dogs []dogrepo.Dog
	err = json.Unmarshal(data, &dogs)
	if err != nil {
		t.Fatal(err)
	}
}

func setUpMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		repository := dogrepo.DogFosterRepository{}

		dog := repository.ExampleDog()

		data, _ := json.Marshal(&dog)

		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	}))
}
