package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tjheslin1/ProfileGoRest/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	go server.Start(logger)

	var input string
	fmt.Scanln(&input)
}
