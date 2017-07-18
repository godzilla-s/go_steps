//这里使用的是官方的JSON, https://github.com/bitly/go-simplejson第三方JSON这个包推荐用
package main

import (
	"fmt"
	"encoding/json"
	"os"
)

//结构体转JSON时，需要定义JSON中相应的标签
type Config struct {
	Host		string		`json:"host"`
	Port		int			`json:"port"`
	FilePath	string		`json:"file_path"`	
	Log			string		`json:"log"`
}

type UserInfo struct {
	userid		int			`json:"userId"`
	username	string		`json:"userName"`
	favour		[]string	`json:"favour"`
	item		struct {
					name	string		`json:"Name"`
					attr	string		`json:"Attr"`
				} `json:"Item"`
}

func main() {
	js := `{"host":"192.168.1.46", "port":8001, "file_path":"/home/hdp/log", "log": "Test"}`
	
	/* JSON 转 struct */
	var cf Config
	if err := json.Unmarshal([]byte(js), &cf); err == nil {
		fmt.Println(cf)
		fmt.Println("Host", cf.Host, "Port:", cf.Port)
	}

	/* JSON 转 map */
	var cfm map[string]interface{}
	if err := json.Unmarshal([]byte(js), &cfm); err == nil {
		fmt.Println(cfm)
		fmt.Println("Host:", cfm["host"], "Port:", cfm["port"])
	}

	/* struct 转 JSON */
	var cf3 = Config{"192.168.1.33", 8005, "home/ujk/src", "Buffer.log"}
	if bs, err := json.Marshal(cf3); err == nil {
		fmt.Println("Struct to JSON: ", string(bs))
	}

	config := Config {
		Host: "192.168.92.130",
		Port: 8201,
		FilePath: "/home/to/path",
		Log: "TLOG.date.log",
	}	

	//转化JSON
	bs, _:= json.Marshal(config)
	os.Stdout.Write(bs) //输出
	println()

	a := make(map[string]interface{})
	a["CarName"] = "Audi A6L"
	a["Price"] = 86.7
	a["Produce"] = "German"
	
	buf, _ := json.Marshal(a)
	fmt.Println(string(buf))
}

