// 基本变量使用
package main

import "fmt"

func test() {
	s := `a
		b\r\n\x00
		c`
	s1 := "Hello, " + 
		"I am Newer"

	/*  下面写法错误， + 不能换行写
	s2 := "Hello, "
		 + "I just make a error"
	*/

	println(s)
	println(s1)
}

//rune 类型变量使用 
func rune_test(){
	fmt.Printf("%T\n", 'a')

	var c1, c2 rune = '\u5677', '伟'	
	fmt.Println(c1, c2, string(c1), string(c2))

	//字符转换
	a := "Fuck"
	b := []byte(a)  // string -> byte, b会被重新分配内存
	b[1] = 'U'
	fmt.Println(string(b))

	c := "电脑"
	d := []rune(c) // string -> rune
	d[1] = '话'
	fmt.Println(string(d))
}

func fortest_str() {
	e := "abCd汉子中国Win"
	for i:=0; i<len(e); i++ {
		fmt.Printf("[%c] ", e[i])
	}
	println()

	for _, f := range(e) {
		fmt.Printf("[%c] ", f);
	}
	println()
}

func main() {
	test()
	rune_test()
	fortest_str()

}
