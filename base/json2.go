package main

import (
	"fmt"
	"encoding/json"
	_"os"
)

type Config struct {
	Host		string		`json:"host"`
	Port		int			`json:"port"`
	FilePath	string		`json:"file_path"`	
	Log			string		`json:"log"`
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

	
}
