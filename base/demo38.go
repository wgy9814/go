package main

import (
	"fmt"
)


func main() {
	num1 := 100
	fmt.Printf("num1的类型是%T, num1的值是%v, num1的地址是%v \n ", num1, num1, &num1)

	//表达式new(T)将创建一个T类型的匿名变量，所做的是为T类型的新值分配并清零一块内存空间，主要用来分配值类型
	//然后将这块内存空间的地址作为结果返回，而这个结果就是指向这个新的T类型值的指针值，返回的指针类型为*T。
	num2 := new(int) // *int
	//num2的类型%T =》 *int
	//num2的值%v =》 地址 0x1140a0ec (这个地址是系统分配的)
	//num2的地址%v =》 地址 0x11406148  (这个地址是系统分配的)
	//num2指向的值 = 100
	*num2 = 100
	fmt.Printf("num2的类型是%T, num2的值是%v, num2的地址是%v \n num2这个指针 指向的值= %v",
		num2, num2, &num2, *num2)

}
