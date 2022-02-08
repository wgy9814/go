package main

import (
	"fmt"
	"go_code/chatroom/client/process"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {
	// 接受用户的选择
	var key int

	//var loop = true
	for {
		fmt.Println("登录多人聊天")
		fmt.Println("1 登录")
		fmt.Println("2 注册")
		fmt.Println("3 退出")
		fmt.Println("请选择(1-3)")
		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			// 登录
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入密码:")
			fmt.Scanf("%s\n",&userPwd)

			//完成登录
			//1. 创建一个 UserProcess 实例
			up := &process.UserProcess{}
			up.Login(userId,userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入密码:")
			fmt.Scanf("%s\n",&userPwd)
			fmt.Println("请输入用户名(nickname):")
			fmt.Scanf("%s\n",&userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			// 退出系统
			os.Exit(0)
			//return
		default:
			fmt.Println("输入有误")
		}
	}
	//if key == 1 {
	//	fmt.Println("请输入用户ID:")
	//	fmt.Scanf("%d\n",&userId)
	//	fmt.Println("请输入密码:")
	//	fmt.Scanf("%s\n",&password)
	//	//fmt.Printf("userId = %d password=%s\n", userId, password)
	//
	//	login(userId,password)
	//}
}