package main

import (
	"fmt"
	"net"
)

func listen() {
	server_address := "0.0.0.0:6969"
	listener, err := net.Listen("tcp", server_address)
	if err != nil {
		panic("couldn't connect with " + server_address)
	}

	fmt.Printf("listening to %s\n", server_address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("invalid connection!")
		}
		fmt.Println("accepted tcp connection from " + conn.RemoteAddr().String())
	}
}
