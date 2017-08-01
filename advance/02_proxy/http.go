// 实现http的代理
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type HttpProxy struct {
	Proxy
}

func (httpProxy *HttpProxy) Run() {
	log.Println("Http Server is On", httpProxy.proxyAddr)
	http.ListenAndServe(httpProxy.proxyAddr, httpProxy)
}

var transport = &http.Transport {
	Proxy:http.ProxyFromEnvironment,
	ResponseHeaderTimeout:30 * time.Second,
}

func (httpProxy *HttpProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "CONNECT" {
		hij, err := w.(http.Hijacker)
		if !ok {
			log.Println("httpserver does not support hijacker")
			return
		}

		proxyCli, err := hij.Hijack()
		if err != nil {
			log.Println("Cannot hijack connection" + err.Error())
			return
		}

		httpProxy.serve(proxyCli, r)
		return
	}

	r.RequestURL = ""
}
