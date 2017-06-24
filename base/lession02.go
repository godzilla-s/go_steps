package main
 
import "fmt"
import "unsafe"

func test() {
	type data struct {
		a int
	}

	var d= data{ 1234 }
	var p *data

	p = &d

	fmt.Printf("d: %d, p: %p\n", d.a, &d)
	fmt.Printf("p: %d, p: %p\n", p.a, p)

	//p++  error, 不能对指针做加减法
}

// Unsafe Point unsafe.Pointer 和任意类型指针间进行转化
func test_point_return() *int{   //返回一个 int行指针 函数
	a := 100
	return &a
}

func unsafePoint_test() {
	x := 0x12345678
	fmt.Printf("a type: %T, %T\n", x, &x)
	p := unsafe.Pointer(&x)  // *int -> Pointer
	n := (*[4]byte)(p)		 // Pointer -> *[4]byte
	
	for i:=0; i<len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
	println()

	b := test_point_return()
	println(b)
	println(*b)

}

func pointer_uintptr_transfer() {
	c := struct { 
		d string
		e int
	} {"MoonLight", 12}

	f := uintptr(unsafe.Pointer(&c))   // *struct -> Pointer -> uintptr
	fmt.Println("c ptr:", f)
	x := f + unsafe.Offsetof(c.d)
	fmt.Println("d ptr:", x)
	f += unsafe.Offsetof(c.e)    // f -> c.e
	fmt.Println("e ptr:", f)


	g := unsafe.Pointer(f)		//uintptr -> Pointer
	h := (*int)(g)  		// Pointer -> *int

	*h = 30					// 相当于 c.e=30

	fmt.Printf("%#v\n", c)
	
}
func main() {
	test()	

	unsafePoint_test()
	pointer_uintptr_transfer()
}
