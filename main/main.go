package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)

	go startServer()

	var input string
	fmt.Scanln(&input)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	io.WriteString(w, "hello, world!\n")
}

func startServer() {
	err := http.ListenAndServe(":6060", nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
