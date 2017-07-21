Golang 笔记
==========================================

### Golang内置函数
-----------------
* func append(slice []Type, elems ...Type) []Type<br>
	如：将数据添加的slice里面
	var a []byte
	a = append(a, 'A')
* func cap(v Type) int
* func close(c chan<- Type)
* func complex(r, i FloatType) ComplexType
* func copy(dst, src []Type) int
* func delete(m map[Type]Type1, key Type)
* func imag(c ComplexType) FloatType
* func len(v Type) int
* func make(Type, size IntegerType) Type
* func new(Type) *Type
* func panic(v interface{})
* func real(c ComplexType) FloatType
* func recover() interface{}

		
### Golang创建项目
----------------
	example: 创建一个project<br>
	新建 project 文件夹<br>
	并在其下面新建三个文件夹, 如下: <br>
		project
			|--bin
			|--pkg
			|--src

	将project所在目录添加到GOPATH环境变量中去<br>
		export $GOPATH=....

	新的package包中的对外函数API, 首字母需大写<br
