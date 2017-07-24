package main

import (
	"fmt"
	"net"
	"log"
)

func main() {
	lsn, err := net.ListenUDP("udp4", &net.UDPAddr{
				IP: net.IPv4(127,0,0,1),
				Port: 9010,})
	if err != nil {
		log.Fatal("listen err:", err)
	}

	fmt.Println("UDP Server On")
	defer lsn.Close()

	for{
		data := make([]byte, 1028)
		read, remoteAddr, err := lsn.ReadFromUDP(data)
		if err != nil {
			fmt.Println("Read Data Fail:", err)
			continue
		}

		fmt.Println(read, remoteAddr)
		fmt.Printf("Data: %s\n", data)

		send := []byte("HelloClient")
		_, err = lsn.WriteToUDP(send, remoteAddr)
		if err != nil {
			fmt.Println("Send Data Fail:", err)
		}
	}
}
