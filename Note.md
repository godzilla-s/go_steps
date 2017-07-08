###########################################
#	Golang 笔记
###########################################

# 1. 创建golang项目
	example: 创建一个project
	新建 project 文件夹
	并在其下面新建三个文件夹, 如下
		project
			|--bin
			|--pkg
			|--src

	将project所在目录添加到GOPATH环境变量中去
		export $GOPATH=....

#2. 新的package包中的对外函数API, 首字母需大写
