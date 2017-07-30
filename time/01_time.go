package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("类型: %T, %s\n", now, now)

	fmt.Println("time.Unix(): ", now.Unix())

	//创建本地时间: 参数sec, usec是从1970-01-01 08:00:00所经过的秒数和纳秒数
	fmt.Println("time.Unix(10, 5000): ", time.Unix(10, 5000))

	fmt.Println("time.Local(): ", now.Local())
	
	//now 是一个time.Time类型
	fmt.Printf("time.Time:\n")
	
	fmt.Printf("Year:%d, Month:%d, Day:%d\n", now.Year(), now.Month(), now.Day())
	
	fmt.Printf("今天是一年中的第 %d 天\n", now.YearDay())
	
	fmt.Printf("现在是: %d 时 %d 分 %d 秒\n", now.Hour(), now.Minute(), now.Second())
	
	fmt.Printf("今天星期 %d\n", now.Weekday())
	
	fmt.Printf("Nano秒: %d\n", now.Nanosecond())

	//Date()函数
	year, mon, day := now.Date()
	fmt.Println(year, mon, day)	
	hour, min, sec := now.Clock()
	fmt.Println(hour, min, sec)

	//时区
	zone, offset := now.Zone()
	fmt.Println("时区: ", zone, offset)
}
