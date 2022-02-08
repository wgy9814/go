package main

import "fmt"

func main() {

	//从键盘输入个数不确定的整数 并判断输入的正数和负数个数 输入为0时结束程序

	var positiveCount int //正数的个数
	var negetiveCount int //负数的个数
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num == 0 {
			break//终止循环
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negetiveCount++

	}
	fmt.Printf("正数的个数是%v 负数的个数是%v \n", positiveCount, negetiveCount )
}