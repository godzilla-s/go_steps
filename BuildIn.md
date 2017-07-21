Golang 内置函数
================

		append 	: 将数据加到slice里面，返回修改后的slice
		>> var a []byte
		>> a = append(a, 'a')
		close 	: 关闭channel
		delete 	: 从map里删除key对应的键值数据
		panic 	: 停止常规的goroutine
		recover	: 允许程序定义goroutine的panic动作
		imag		: 返回complex的实部
		real		: 返回complex的虚部
		make		: 创建一个类型数据(只能应用于slice, map, channel)
		new		  : 创建一个类型数据
		cap		  :
		copy		  :
		len		  :
