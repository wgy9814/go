package main

import (
	"errors"
	"fmt"
)

//默认情况下 当发生错误后（panic）,程序就会退出
func test(){
	//使用defer + recove 来捕获和处理异常
	defer func() {
		err := recover() //recover()内置函数，可以捕获到异常
		if err != nil { //说明捕获到错误
			fmt.Println("err = ", err)
			//这里可以将错误信息发送给管理员
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取..
		return nil
	} else {
		return errors.New("读取文件错误....")
	}
}
func test02() {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行....")
}
func main() {
	//测试
	//test()
	//fmt.Println("main()下面的代码....")

	//测试自定义错误的使用
	test02()
	fmt.Println("main()下面的代码....")
}
