package main

import (
	"fmt"
)

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

	fmt.Println("length a:", len(a), "capacity:", cap(a))
	b := a[2:4]
	fmt.Println(b)

	c := []int{23, 45, 67, 12, 56, 78, 30, 53, 47}
	//注意切片end位置
	fmt.Println(c[:2], c[1:3], c[2:6], c[5:])
	
	d := []string{"John", "Boby", "Lucia", "罗罗", "张晓波"}
	//注意d[0] 返回是字符串, d[1:2]返回切片，数组下标不支持负数,
	//d[:]
	fmt.Println(d[0], d[1:2], d[:])

	//切片遍历
	for _, v := range d {
		fmt.Printf("%s ", v)
	}
	println()
	
	for i:=0; i<len(d); i++ {
		fmt.Printf("%s ", d[i])
	}
	println()

	//使用make
	e := make([]int, 3)
	e[0] = 23
	e[1] = 35
	e[2] = 45
	fmt.Println(e)
	//e[3] = 50 	//下标越界

	//len, cap
	f := make([]string, 3, 5)
	f[0] = "心比天高"
	f[1] = "脚踏实地"
	f[2] = "笑脸敞开"
	fmt.Printf("addr: %p, %s\n", &f, f)
	//f[3] = "勇往直前" //下标越界
	f = append(f, "勇往直前")
	fmt.Printf("addr: %p, %s\n", &f, f)
	//使用append可以继续添加, 而cap也会自动增加, 
	//其增加有规律
	f = append(f, "前赴后继")
	f = append(f, "是否继续")
	f = append(f, "再加一个")
	fmt.Printf("addr: %p, %s\n", &f, f)
	fmt.Println("afper append: length-", len(f), "cap-", cap(f))
	//slice 合并
	h := []int{1, 2, 3, 4}
	g := []int{5, 6, 7, 8}
	h = append(h, g...)
	fmt.Println("合并:", h)

	/*	不同类型不能append
	j := []string{"ofcourse", "smile"}
	h = append(h, j)
	fmt.Println("能否添加字符串:", h)
	*/

	h = append(h[:5], g[2:]...)
	fmt.Println(h)
}
