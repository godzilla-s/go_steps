package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"io/ioutil"
)

func ReadInputByDelim() {
	reader := bufio.NewReader(os.Stdin)
	//line, err := reader.ReadString('\n')
	//line, _ := reader.ReadBytes(' ')
	line, _ := reader.ReadSlice(' ')
	fmt.Println(string(line))
}

func ReadInputByLine() {
	reader := bufio.NewReader(os.Stdin)
	//line, isPrefix, _ := reader.ReadLine()
	//fmt.Println(string(line), isPrefix)
	r, size, _ := reader.ReadRune()
	fmt.Println(string(r), size)
}

func ReadFileByDelim() {
	f, err := os.Open("Text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	/*reader : way - 01*/
	rbuf := make([]byte, 256)
	len, _ := f.Read(rbuf)
	fmt.Printf("file read: %d\n%s", len, string(rbuf))

	/*reader : way - 02*/
	f.Seek(0, os.SEEK_SET)	
	reader := bufio.NewReader(f)  /*return a Reader object*/
	fmt.Println("bufio reader:")
	for {
		//line, rerr := reader.ReadString('\n')
		line, rerr := reader.ReadBytes('\n')
		if rerr == io.EOF {
			break
		}
		fmt.Println(line)
		//fmt.Printf(string(line))
	}
	
	/*reader : way - 03*/
	f.Seek(0, os.SEEK_SET)
	buf, _ := ioutil.ReadAll(f)
	fmt.Printf("ioutil read:\n%s", buf)
}

func main(){
	//ReadInputByLine()
	//ReadInputByDelim()
	ReadFileByDelim()

}
