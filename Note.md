	Golang 笔记
==========================================

***Golang内置函数
	append 	: 将数据加到slice里面，返回修改后的slice
	close 	: 关闭channel
	delete 	: 从map里删除key对应的键值数据
	panic 	: 停止常规的goroutine
	recover	: 允许程序定义goroutine的panic动作
	imag	: 返回complex的实部
	real	: 返回complex的虚部
	make	: 创建一个类型数据(只能应用于slice, map, channel)
	new	: 创建一个类型数据
	capy	:
	copy	:
	len	:
		
***Golang创建项目
	example: 创建一个project
	新建 project 文件夹
	并在其下面新建三个文件夹, 如下
		project
			|--bin
			|--pkg
			|--src

	将project所在目录添加到GOPATH环境变量中去
		export $GOPATH=....

	新的package包中的对外函数API, 首字母需大写
