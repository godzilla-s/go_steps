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
	conn.Read(buffer)

	fmt.Println(string(buffer))

	for {
		fmt.Print(">:")	
		fmt.Scanf("%s", &buffer)
		
		if string(buffer) == "quit" {
			os.Exit(1)
		}

		conn.Write(buffer)
		time.Sleep(100 * time.Millisecond)
	}
	
}
