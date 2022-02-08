package main

import "fmt"

func main() {

	//方式二
	//1.演示切片的基本使用 make
	//通过make方式创建的切片对应的数组是由make底层维护，对外不可见，只能通过slice去访问各个元素
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[2] = 20
	//对于切片 必须make使用
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))

	//第3种方式：定义一个切片，直接就指定具体数值，使用原理类似make的方式
	var strSlice []string = []string{"tom","jack","mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice的size=", len(strSlice))
	fmt.Println("strSlice的cap=", cap(strSlice))
}
