package main

import "fmt"

func test() int {
	return 90
}
func main() {

	//赋值运算符的使用演示
	//var i int
	//i = 10 //基本赋值

	//有两个变量，a和b，要求将其交换，打印结果
	//a := 9
	//b := 2
	//fmt.Printf("交换前的值是 a= %v,b=%v \n",a,b)
	//t := a
	//a = b
	//b = t
	//fmt.Printf("交换后的值是 a= %v,b=%v \n",a,b)
	//
	//
	////复合赋值的操作
	//a += 17 //等价a = a+17
	//
	//fmt.Println("a=",a)
	//
	//var c int
	//c = a + 3 //赋值运算的执行顺序是从右到左
	//fmt.Println("c=",c)
	//
	////赋值运算符的左边 只能是变量，右边可以是变量，表达式，常量值
	////表达式：任何有值的都可以看做是表达式
	//var d int
	//d = a
	//d = 8 + 2 * 8
	//d = test() + 90
	//d = 890
	//fmt.Println("d=",d)

	//有两个变量 a和b 要求将其交换 但是不允许使用中间变量 最终打印结果
	var a int = 10
	var b int =20
	a = a + b//
	b = a - b// b = a + b - b ==> b = a
	a = a - b// a = a + b - a ==> a = a
	fmt.Printf("a=%v b=%v",a,b)
}