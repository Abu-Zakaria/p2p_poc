package main

import (
	"fmt"
	"net"
)

func connect() {
	friend_address := "103.121.104.177:6969"

	conn, err := net.Dial("tcp", friend_address)
	if err != nil {
		panic("couldn't connect with " + friend_address)
	}

	fmt.Printf("connected with to %s\n", friend_address)
	fmt.Println(conn)
}
