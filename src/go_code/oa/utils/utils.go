package utils

import "fmt"

//将计算的功能，放到一个函数里面，调用即可
//函数首字母大写 类似其他语言的public
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("输入错误")
	}
	return res
}