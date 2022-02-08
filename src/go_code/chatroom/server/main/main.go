package main

import (
	"fmt"
	"go_code/chatroom/server/model"
	"net"
	"time"
)


//处理核客户端的通讯
func process(conn net.Conn){
	// 延时关闭conn
	defer conn.Close()

	//这里调用总控。创建一个
	processor := &Processor{
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}
func init() {
	//当服务器启动时，我们就去初始化我们的redis连接池
	initPool("localhost:6379", 16, 0, 300 * time.Second);
	initUserDao()
}
//我们编写一个函数，完成对UserDao的初始化工作
func initUserDao(){
	//这里的pool是全局变量
	//这里需要注意一个初始化顺序问题
	//initPool 在  initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {

	fmt.Println("服务器[新的结构]在8889端口监听")
	listener,err := net.Listen("tcp","127.0.0.1:8889")
	defer listener.Close()
	if err != nil{
		fmt.Println("监听错误",err)
		return
	}
	// 一旦监听成功，等待客户端链接服务器
	for  {
		fmt.Println("等待链接")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("链接出错",err)
		}
		// 链接成功,启动协程和客户端保持通信
		go process(conn)

	}
}