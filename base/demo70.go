package main

import "fmt"

//结构体
type Circle struct {
	radius float64
}


//声明一个方法 area 和 Circle绑定，可以返回面积
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

//为了提高效率 通常我们方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	//因为c是指针，因此我们标准的访问其字段的方式是 (*c).radius
	//return 3.14 * (*c).radius * (*c).radius
	c.radius = 10.0
	return 3.14 * c.radius * c.radius
}





func main() {
	//声明有结构体Circle，字段为radius
	//声明一个方法 area和 Circle绑定，可以返回面积
	//画出area执行过程

	//创建一个 Circle 变量
	//var c Circle
	//c.radius = 4.0
	//res := c.area()
	//fmt.Println("面积=", res)

	//创建一个 Circle 变量
	var c Circle
	c.radius = 5.0
	res2 := (&c).area2()
	//编译器底层做了优化 (&c).area() 等价于 c.area()
	fmt.Println("面积=", res2)
	fmt.Println("c.radius=", c.radius)
}
