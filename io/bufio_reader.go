/* bufio 包实现有缓冲的IO读写 */
package main

import (
	"fmt"
	"bufio"
	"strings"
	"log"
	"io"
)

func main() {
	//创建一个bufio Reader 
	buf1 := strings.NewReader("Hello, Dear World, 完美世界")
	reader1 := bufio.NewReader(buf1)

	fmt.Println("可读字节数: ", reader1.Buffered())
	buffer, err := reader1.Peek(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("使用Peek:", string(buffer), "可读字节数:", reader1.Buffered())

	//使用ReadByte 一个字节一个字节读取
	ch1, _ := reader1.ReadByte()
	fmt.Println(ch1, "可读字节数:", reader1.Buffered())
	ch2, _ := reader1.ReadByte()
	fmt.Println(ch2, "可读字节数:", reader1.Buffered())
	for {
		ch, err := reader1.ReadByte()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}

		fmt.Print(ch, " ")
	}
	println("可读字节数:", reader1.Buffered()) // 全部读取完之后，buffered为0

	buf2 := strings.NewReader("Be Best To Be a Millionaire")
	reader2 := bufio.NewReader(buf2)
	//使用ReadBytes: 读取字符知道遇到'd'
	data2, err := reader2.ReadBytes('d') 
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println("使用ReadBytes: ", string(data2), "可读字节数:", reader2.Buffered())
	
	buf3 := strings.NewReader("Rain and Wind are together, it's cold")	
	reader3 := bufio.NewReader(buf3)
	data3, err := reader3.ReadSlice('g')	
	if  err != nil && err != io.EOF {
		log.Fatal(err)
	}		
	fmt.Println(string(data3), "可读字节数:", reader3.Buffered())

	//ReadRune：读取一个utf-8编码的unicode码值
	buf4 := strings.NewReader("怎么一直老下雨，好久没见太阳了")
	reader4 := bufio.NewReader(buf4)
	data4, _, _ := reader4.ReadRune()
	fmt.Println(string(data4), "可读字节数:", reader4.Buffered())
}
