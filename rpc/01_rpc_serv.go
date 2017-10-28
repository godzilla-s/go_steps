package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

type Handle int 

const step = 2

func (h *Handle) GetInfo(arg int, result *int) error {
	*result = arg + step
	return nil 
}

type Member struct {
	id  int
	name string 
}

var membase = make(map[int]Member)


func main() {
	http.HandleFunc("/hello", helloTest)

	handle := new(Handle)
	rpc.Register(handle)
	rpc.HandleHTTP()

	lsn, err := net.Listen("tcp", ":3003")
	if err != nil {
		fmt.Println("fail to listen: ", err)
		return 
	}

	fmt.Println("Listen Port OK")
	http.Serve(lsn, nil)
}

func helloTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>hello your name</body></html>")
}
