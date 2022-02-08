package main

import (
	"fmt"
	"strings"
)

//闭包就是一个函数和与其相关的引用环境组合的一个整体（实体）
//累加器
//1.AddUpper是一个函数，返回的数据类型是 func (int) int
//2.闭包是类，函数是操作，n是字段，函数和他使用到n构成闭包
//3.当我们反复的调用f函数时，因为n是初始化一次，因此每调用一次就累加	
func AddUpper() func (int) int {
	//返回的事一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36) //int转换成ascii码 36=》‘$’
		fmt.Println("str=",str)
		return n
	}
}
//1）编写一个函数 makesuffix（suffix string）可以接收一个文件后缀名（比如．jpg），并返回一个闭包
//2）调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如．jpg），则返回 文件名．jpg，如果已经有.jpg后缀，则返回原文件名
//3）要求使用闭包的方式完成
//4）strings.HasSuffix，该函数可以判断某个字符串是否有指定的后缀。

//1）返回的匿名函数和 makeSuffix （suffix string）的suffix变量组合成一个闭包，因为返回的函数引
//用到suffix这个变量
//（2）我们体会一下闭包的好处，如果使用传统的方法，也可以轻松实现这个功能，但是传统方法
//需要每次都传入后缀名，比如.jpg，而闭包因为可以保留上次引用的某个值，所以我们传入
//一次就可以反复使用。大家可以仔细的体会一把！

func makeSuffix(suffix string) func (string) string {
	return func (name string) string {
		//如果name没有指定的后缀，则加上否则返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {

	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	//返回一个背包
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=",f2("winter"))
	fmt.Println("文件名处理后=",f2("fly.jpg"))

}

