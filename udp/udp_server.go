package udp

import (
	"fmt"
	"log"
	"net"
)

func CreateUDPServer() {
	server_host := "0.0.0.0"
	server_port := ":5895"

	udpAddr, err := net.ResolveUDPAddr("udp4", server_host+server_port)
	if err != nil {
		panic(err)
	}
	fmt.Println("resolved udp address:", udpAddr)

	conn, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("listening to", server_host+server_port)

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		size, addr, err := conn.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		log.Println("size:", size)
		log.Println("addr:", addr)
		log.Println("buffer:", string(buf))
	}
}
