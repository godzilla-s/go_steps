package main

import (
	"fmt"
	"time"
)

//Duration类型代表两个时间点之间经过的时间，以纳秒为单位。
//可表示的最长时间段大约290年。
//Duration 类型为 int64

func main() {
	//将下面数据转化为以second为单位的时间
	now := 100
	fmt.Println("100s : ", time.Duration(now) * time.Second)
	
	pass := 100000
	fmt.Println("100000s : ", time.Duration(pass) * time.Second)	

	//将时间转化为Duration
	d, err := time.ParseDuration("10h20m12s")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("Duartion: %T, %v, ToString: %s\n", d, d, d.String())
	//转化为小时为单数，分为单位，秒为单位
	fmt.Println("Hour:", d.Hours(), "Minute:", d.Minutes(), "Second:", d.Seconds())
	fmt.Println("NanoSecond:", d.Nanoseconds())
}
