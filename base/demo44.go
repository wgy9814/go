package main

import "fmt"

func main() {
	//方式1
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明/定义一个切片
	//slice := intArr[1:3]
	//1.slice就是切片名
	//2.intArr[1:3] 表示slice 引用到 intArr 这个数组
	//3.引用intArr数组起始下标为1，最后下标为3（但是不包含3）
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素=", slice) //22 33
	fmt.Println("slice 的元素个数=", len(slice)) //2
	fmt.Println("slice 的容量=", cap(slice))//切片的容量是可以动态变化的

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0==%v]\n", &slice[0], slice[0])
	slice[1] = 34
	fmt.Println()
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素=", slice)

	//slice地址取的是起始位置数组的地址
	//slice是一个引用类型
	//slice从底层来讲，其实就是一个数据结构（struct结构体）
//	type slice struct {
//		ptr *[2]int
//		len int
//		cap int
//	}
}
