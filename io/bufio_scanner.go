package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	fr, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//扫描读取文件(一行一行)
	scanner := bufio.NewScanner(fr)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	scanner2 := bufio.NewScanner(fr)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 获取一个单词
		advance, token, err = bufio.ScanWords(data, atEOF)
		// 判断其能否转换为整数，如果不能则返回错误
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		// 这里包含了 return 0, nil, nil 的情况
		return
	}

	scanner2.Split(split) //必须在Scan()函数前
	for scanner2.Scan() {
		line := scanner2.Text()
		fmt.Println(line)
	}

	//使用os.Stdio: 输入数据
	scanner3 := bufio.NewScanner(os.Stdin)
	for scanner3.Scan() {
		line := scanner3.Text()
		if line == "quit" {
			break
		}
		fmt.Println(scanner3.Text(), "\t[type 'quit' 退出]")
	}

	if err := scanner3.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//切分单词
	const text = "When sun is rising, new hope is around us, birds sings, people begin working"
	scanner4 := bufio.NewScanner(strings.NewReader(text))
	scanner4.Split(bufio.ScanWords)
	for scanner4.Scan() {
		fmt.Println(scanner4.Text())
	}
}
