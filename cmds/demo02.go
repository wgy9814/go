package main

import (
	"flag"
	"fmt"
)

func main() {

	// 定义几个变量，用于接收命令行的参数值
	var user        string
	var pwd    string
	var host        string
	var port        int

	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	//"u",就是 -u指定参数
	//"",默认值
	//"用户名，默认为空" 说明
	flag.StringVar(&user, "u", "root", "账号，默认为root")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	// 解析命令行参数写入注册的flag里
	//这里有一个非常重要的操作，转换，必须调用该方法
	flag.Parse()
	// 输出结果
	fmt.Printf("user：%v\npassword：%v\nhost：%v\nport：%v\n",
		user, pwd, host, port)



}