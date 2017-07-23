//使用get获取网页信息

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	//resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	resp, err := http.Get("http://www.ltaaa.com/")
	if err != nil {
		fmt.Println(err)
		return
	}	

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
