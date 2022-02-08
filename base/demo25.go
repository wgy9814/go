package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}
func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}

func main() {
	//递归调用
	//test(4)
	test2(4)
}