// 获取公网地址
package main

import (
	"regexp"
	"fmt"
	"time"
	"net"
	"net/http"
	"io/ioutil"
)

// Linux : 可以通过命令: echo `nc ns1.dnspod.net 6666`

// 方法一:
func getPublicIP_01() {
	conn, err := net.DialTimeout("tcp", "ns1.dnspod.net:6666", 100 * time.Second)
	if err != nil {
		fmt.Println("fail to dail:", err)
		return 
	}

	defer conn.Close() 

	var bytes []byte  
	deadline := time.Now().Add(5 * time.Second)  
	err = conn.SetDeadline(deadline)  
	if err != nil {  
		fmt.Println("fail to set deadline:", err)
		return 
	} 

	bytes, err = ioutil.ReadAll(conn)  
	if err != nil {  
		fmt.Println("fail to read:", err )
		return 
	}

	fmt.Println("PublicIP:", string(bytes))
}

// 方法二: 通过获取网页地址解析
func getPublicIP_02(netaddr string) {
	res, err := http.Get(netaddr)
	if err != nil {
		fmt.Println("fail to get:", err)
		return 
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read all:", err)
		return 
	}
	reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	fmt.Println("PublicIP:", reg.FindString(string(result)))
}

func main() {
	getPublicIP_01()
	getPublicIP_02("http://iframe.ip138.com/ic.asp")
	getPublicIP_02("http://myexternalip.com/raw")
}