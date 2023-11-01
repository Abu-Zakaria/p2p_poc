package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("not enough arguments")
	}

	arg1 := args[1]

	if arg1 == "listen" {
		listen()
	} else if arg1 == "connect" {
		connect()
	} else {
		fmt.Println("wat da hell u want buddy?")
	}
}
