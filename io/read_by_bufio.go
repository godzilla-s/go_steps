package main

import(
	"fmt"
	"bufio"
	"os"
	"io"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}	

	defer file.Close()

	reader := bufio.NewReader(file)
	fmt.Println("Buffered: ", reader.Buffered())

	fmt.Println("ReadByte:")
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("finish")
			break
		}
		print(string(line))
	}

	reader.Reset(file)
	fmt.Println("ReadLine:")
	for {
		/* 一行一行读取, 不建议使用 */
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println("read finish")
			break	
		}
		println(string(line))	
	}

	fmt.Println("ReadString:")
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Println(line)
	}

	fmt.Println("NewScanner:")
	scanner := bufio.NewScanner(file)
	fmt.Println(scanner.Text())
}
