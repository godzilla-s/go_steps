package main

import (
	"fmt"
	"net"
	"log"
)

func main() {
	conn, err := net.DialUDP("udp4", nil,  &net.UDPAddr{
				IP: net.IPv4(127,0,0,1),
				Port: 9010})
	if err != nil {
		log.Fatal("Dail err:", err)
	}

	defer conn.Close()

	send := []byte("Hello Server")
	_, err = conn.Write(send)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	data := make([]byte, 512)
	read, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(read, remoteAddr)
	fmt.Printf("recv: %s\n", data)
}
