package main

import (
	"fmt"
	"flag"
	"net"
	"os"
	"time"
	"bufio"
	"encoding/json"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "9999", "port")

type Message struct {
	Buffer  string 	`json:"data"`
	Type 	int 	`json:"type"`
}

type Response struct {
	Buffer	string	`json:"data"`
	Status 	int 	`json:"status"`
}

func main() {
	flag.Parse()  /* 解析参数 */
	lsn, err := net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Listen err:", err)
		os.Exit(1)
	}
	defer lsn.Close()

	fmt.Println("Listen is Start")

	for {
		cli, err := lsn.Accept()
		if err != nil {
			fmt.Println("Accept err: ", err)
			os.Exit(1)
		}

		fmt.Println("New client: ", cli.RemoteAddr())
		go handleConn(cli)
	}
}

func handleConn(c net.Conn) {
	remoteIP := c.RemoteAddr().String()
	defer func(){
		fmt.Println("Disconnected: ", remoteIP)
		c.Close()
	}()

	//初始化读写缓存区
	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)

	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			return
		}

		var msg Message
		json.Unmarshal(b, &msg) //反序列化
		fmt.Println("Read : ", msg.Buffer, "Type: ", msg.Type)

		//回复客户端
		response := Response {
			Buffer: time.Now().String(),
			Status: 200,
		}

		r, _ := json.Marshal(response)

		writer.Write(r)
		writer.Write([]byte("\n"))
        writer.Flush()
	}
	fmt.Println("Done!")
}
