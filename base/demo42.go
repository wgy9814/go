package main

import "fmt"

func test01(arr [3]int) {
	arr[0] = 233
}

func test02(arr *[3]int) {
	fmt.Printf("arr指针的地址=%p", &arr)
	(*arr)[0] = 233
}
func main() {

	//数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 30
	//这里会报错 类型不一样
	//arr01[2] = 1.1
	//长度是固定的，不能动态变化,否则报越界
	//arr01[3] = 890
	//fmt.Println(arr01)

	//数组创建后，如果没有赋值，有默认值（零值）
	//1. 数值（整数类型，浮点数类型） =》 0
	//2. 字符串 ==》 ""
	//3. 布尔类型 ==》 false
	var arrs01 [3]float32
	var arrs02 [3]string
	var arrs03 [3]bool
	fmt.Printf("arrs01=%v, arrs02=%v, arrs03=%v\n",arrs01, arrs02, arrs03)

	//go数组属于值类型 在默认情况下是值传递 因此会进行值拷贝 数组间不会相互影响
	arr := [3]int{11, 22, 33}
	test01(arr)
	fmt.Println(arr)
	test02(&arr)
	fmt.Println(arr)

	//长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度
}
