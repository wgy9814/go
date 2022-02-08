package main

import "fmt"

func test(char byte) byte {
	return char + 1;
}
func main() {
	//接收一个字符 a b c d e f g
	//表示星期一 以此类推
	//使用switch完成


	var key byte

	fmt.Println("请输入一个字符，abcdefg")
	fmt.Scanf("%c", &key)
	switch test(key) {
		case 'a':
			fmt.Println("星期一")
		case 'b':
			fmt.Println("星期二")
		case 'c':
			fmt.Println("星期三")
		default:
			fmt.Println("输入有误")
	}
	var n1 int32 = 5
	var n2 int32 = 20
	var n3 int32 = 5
	switch n1 {
		case n2 , 10 , 5:
			fmt.Println("ok1")
		case n3 :
			fmt.Println("ok2")
	}

	//switch也可以不带表达式
	var age int = 10
	switch {
		case age == 10 :
			fmt.Println("age=10")
		case age == 20 :
			fmt.Println("age=20")
		case age == 30 :
			fmt.Println("age=30")

	}

	//case也可以对 范围判断
	var scroe int = 30
	switch {
		case scroe > 90 :
			fmt.Println("成绩优秀")
		case scroe >= 70 && scroe <= 90 :
			fmt.Println("成绩优良")
		case scroe == 30 :
			fmt.Println("成绩不及格")

	}

	//switch后可以声明/定义一个变量，分号结束，不推荐
	switch grade := 90; {
		case grade > 90 :
			fmt.Println("成绩优秀")
		case grade >= 70 && grade <= 90 :
			fmt.Println("成绩优良")
		case grade == 30 :
			fmt.Println("age=30")

	}

	//switch 的穿透 fallthrough
	var num int = 10
	switch num {
		case 10:
			fmt.Println("ok1")
			fallthrough //默认只穿透一层
		case 20:
			fmt.Println("ok2")
		case 30:
			fmt.Println("ok3")
		default:
			fmt.Println("没有匹配到....")
	}

}