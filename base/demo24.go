package main

import (
	"fmt"
	"go_code/oa/utils"
	//util "go_code/oa/utils" //取别名
)

//计算两个数的和和差，并返回结果
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '+'
	result := utils.Cal(n1, n2, operator)
	fmt.Println("result=",result)

	n1 = 4.5
	n2 = 3.6
	operator = '*'
	result2 := utils.Cal(n1, n2, operator)
	fmt.Println("result2=",result2)

	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1=%v,res2=%v \n", res1, res2)

	//希望忽略某个返回值 则使用 _符号表示占位忽略
	_, res3 := getSumAndSub(1, 2)
	fmt.Println("res3=", res3)
}