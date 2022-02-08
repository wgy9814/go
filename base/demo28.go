package main

import "fmt"


//n1就是 *int类型
func test(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test() n1= ", *n1)
}



func main() {
	num := 20
	test(&num)
	fmt.Println("main() num= ", num)
}

