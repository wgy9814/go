package main

import "fmt"

/*
请使用递归的方法，求出斐波那契数列1,1,2，3,5,8,13...
给你一个整数n 求出他的斐波那契数是多少
 */
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	}else{
		return fbn(n - 1) + fbn(n - 2)
	}
}

/*
求函数值 已知 f(1)=3; f(n) = 2*f(n-1)+1;请使用递归 求出f(n)的值
*/
func f(n int) int {
	if n == 1 {
		return 3
	}else{
		return 2 * f(n - 1) + 1
	}
}
func main() {
	//res := fbn()
	//递归调用
	//fmt.Println("res=",fbn(3))
	//fmt.Println("res=",fbn(4))
	//fmt.Println("res=",fbn(5))
	//fmt.Println("res=",fbn(6))

	fmt.Println("f(1)=",f(1))
	fmt.Println("f(5)=",f(5))
}

