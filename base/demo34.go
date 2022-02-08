package main

import (
	"fmt"
)
var Age int = 20
//Name := "tom" //赋值语句不能在函数体外 会报错

//函数的defer
//为什么需要 defer
//在函数中，程序员经常需要创建资源（比如：数据库连接、文件句柄、锁等），为了在函数执行完
//毕后，及时的释放资源，Go的设计者提供defer（延时机制）。
//在defer将语句放入栈时，也会将相关的值拷贝同时入栈

func sum(n1 int, n2 int) int {
	//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer）
	//当函数执行完毕后，再从defer栈，按照先进后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1)//defer 3. ok1 n1=10
	defer fmt.Println("ok2 n2=", n2)//defer 2. ok2 n2=20

	//增加一句话
	n1++
	n2++
	res := n1 + n2 //res = 30
	fmt.Println("ok3 res=", res) //1. ok3 res=30
	return res

}
func main() {
	res := sum(10, 20)
	fmt.Println("res=", res) //4. res=30

}

