package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8010")
	if err != nil {
		fmt.Println("Dail err: ", err)
		os.Exit(1)
	}

	defer conn.Close()

	buffer := make([]byte, 256)
	for {
		fmt.Print(">:")	
		fmt.Scanf("%s", &buffer)
		
		if buffer == "quit" {
			os.Exit(1)
		}

		conn.Write(buffer)
		time.Sleep(100 * time.Millisecond)
	}
}
