package main

import (
	_"fmt"
	"flag"
)

type proxy interface {
	Run()
}

var (
	proxyAddr	string
	backAddr	string
)

func main() {
	flag.StringVar(&proxyAddr, "a", ":9001", "proxy host to listen on")
	flag.StringVar(&backAddr, "b", ":8001", "back host to listen on")
	
	flag.Parse()
	
	var p proxy

	p = &Proxy{proxyAddr: proxyAddr, backAddr: backAddr}

	p.Run()
}

