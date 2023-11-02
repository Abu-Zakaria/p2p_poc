package main

import (
	"github.com/Abu-Zakaria/p2p_poc/udp"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		log.Fatal("args missing!")
	}

	firstArgs := args[1]

	if firstArgs == "udp_server" {
		udp.CreateUDPServer()
	} else {

	}
}
