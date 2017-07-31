package main

import (
	"log"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Println("ERROR 1:", err)
		return
	}
	
	response, err := client.Do(request)
	if err != nil {
		log.Println("ERROR 2:", err)
		return 
	}

	log.Println(response.Header)
	
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ioutil err: ", err)
		return
	}

	log.Println("数据长度: ", len(data))
}
