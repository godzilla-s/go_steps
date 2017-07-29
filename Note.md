# Golang 笔记
==========================

### Go 内置函数
-----------------
* func append(slice []Type, elems ...Type) []Type<br>
	如：将数据添加的slice里面
```go
	var a []byte
	a = append(a, 'A')
```
* func cap(v Type) int
```go
	var a []byte
	fmt.Println(cap(a))
```
* func close(c chan<- Type)
	<br>关闭channel管道
* func complex(r, i FloatType) ComplexType
	<br>复数
* func copy(dst, src []Type) int
	<br>拷贝	
* func delete(m map[Type]Type1, key Type)
	<br>删除map中的数据
```go
	a := make(map[string]string)
	a["Id"] = "123456"
	delete(a, "Id")
```
* func imag(c ComplexType) FloatType
* func len(v Type) int
* func make(Type, size IntegerType) Type
* func new(Type) *Type
* func panic(v interface{})
* func real(c ComplexType) FloatType
* func recover() interface{}

参考： http://lib.csdn.net/article/go/34354

### Go 创建项目
----------------
example: 创建一个project<br>
新建 project 文件夹<br>
并在其下面新建三个文件夹, 如下: <br>
> project
>> |--bin <br>
>> |--pkg <br>
>> |--src <br>

将project所在目录添加到GOPATH环境变量中去<br>
	export $GOPATH=....

新的package包中的对外函数API, 首字母需大写<br>

### Go 命令
-----------------
go run： 可以编译并运行命令源码文件
go test： 用于对Go语言编写的程序进行测试
go build：用于编译我们指定的源码文件或代码包以及它们的依赖包<br>
go install：于编译并安装指定的代码包及它们的依赖包<br>
go get：从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装<br>
go clean: 删除掉执行其它命令时产生的一些文件和目录
go list: 列出指定的代码包的信息


