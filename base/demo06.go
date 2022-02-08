package main

import "fmt"

func test() bool{
	fmt.Println("test.......")
	return true
}
func main() {
	var i int = 10
	//说明 因为i<9为false 所以后面test()就不会执行
	if i < 9 && test() {
		fmt.Println("ok.......")
	}

	//说明 因为i>9为true 所以后面test()就不会执行
	if i > 9 || test() {
		fmt.Println("hello.......")
	}

	//演示关系运算符的使用
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 == n2)
	flag := n1 > n2
	fmt.Println(flag)

	//演示逻辑运算符的操作 &&
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}

	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	//演示逻辑运算符的操作 ||

	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}

	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	//演示逻辑运算符的操作 !

	if age > 30 {
		fmt.Println("ok5")
	}

	if !(age > 30) {
		fmt.Println("ok6")
	}
}