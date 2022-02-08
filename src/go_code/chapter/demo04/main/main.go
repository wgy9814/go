package main

import (
	"fmt"
	"go_code/chapter/demo04/model"
)
//指针类型
func main() {
	//基本数据在内存布局
	var i int = 10
	//i的地址是什么，&i

	fmt.Println("i的地址是",&i)

	//1. ptr是一个指针变量
	//2. ptr的类型 *int
	//3. ptr本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Printf("ptr的地址=%v\n",&ptr)
	fmt.Printf("ptr指向的值=%v\n",*ptr)


	var num int = 9
	fmt.Println("num的地址是",&num)

	var ptx *int
	ptx = &num
	*ptx = 10//这里修改 会影响到num的值变化
	fmt.Println("num=",num)

	var _ab int = 3
	fmt.Println("_ab=",_ab)
	fmt.Println(model.HeroName)



}