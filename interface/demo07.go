package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //ok
	//如何将a赋值给一个 point 变量
	var b Point
	//b = a //error
	b = a.(Point)//类型断言 表示按断a是否指向 Point类型的变量，如果是就转成 Point类型并赋给b变量，否则报错
	fmt.Println(b)

	//类型断言的其他案例
	//var x interface{}
	//var b2 float32 = 1.1
	//x = b2//空接口 可以接收任意类型
	////x=>float32[类型断言],如果类型不匹配，就会报panic，因此进行类型断言时，要确保原来的空接口指向的就是断言的类型
	//y := x.(float32)
	//fmt.Printf("y的类型是 %T 值是%v", y, y)

	//类型断言（带检测）
	var x interface{}
	var b2 float32 = 1.1
	x = b2//空接口 可以接收任意类型
	//x=>float32[类型断言],如果类型不匹配，就会报panic，因此进行类型断言时，要确保原来的空接口指向的就是断言的类型
	//y, ok := x.(float64)
	if y, ok := x.(float64);ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是 %T 值是%v", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")
}
