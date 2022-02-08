package main

import "fmt"

func fbn(n int) ([]uint64) {
	//声明一个切片，切片大小n
	fnbSlice := make([]uint64, n)

	//第一个数和第二个数都是1
	fnbSlice[0] = 1
	fnbSlice[1] = 1
	for i := 2; i < n; i++ {
		fnbSlice[i] = fnbSlice[i - 1] + fnbSlice[i - 2]
	}
	return fnbSlice
}

func main() {

	/*
	1.可以接收一个 n int
	2. 能够将斐波那契数列的数列放到切片中
	3.提示，斐波那契数列的形式：
	arr[0] = 1; arr[1] = 1; arr[2] = 2; arr[3] = 3; arr[4] = 5; arr[5] = 8;

	思路
	1.声明一个函数 fbn(n int) ([]uint64)
	2.编程 fbn(n int)进行for循环存放斐波那契数列的数列
	 */

	fnbSlice := fbn(10)
	fmt.Println("fnbSlice=fnbSlice", fnbSlice)
}
