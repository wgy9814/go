package main

import "fmt"

func main() {

	fmt.Println("Hello, 大叔大婶!")
	var n [10]int
	var i,j int

	for i = 0; i<10;i++ {
		n[i] = i+100
	}
	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}
	a := [3][4]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /* 第三行索引为 2 */
	}
	fmt.Println(a)
	var b int = 10

	fmt.Printf("变量的地址: %x\n", &b  )

	fmt.Println("大叔大婶\r手术"  )
	num := 5.124e2
	fmt.Printf("变量的地址: %T", num  )
	fmt.Println(num)

	var c1 byte = 'a'
	fmt.Println("c1=", c1  )

	fmt.Printf("c1=%c c1=%d", c1,c1)


}

