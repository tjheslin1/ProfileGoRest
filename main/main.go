package main

import (
	"fmt"

	"github.com/tjheslin1/ProfileGoRest/server"
)

func main() {
	go server.Start()

	var input string
	fmt.Scanln(&input)
}
