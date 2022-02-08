package main

import "fmt"

//结构体
type Person struct {
	Name string
}
type Dog struct {

}

//给 Person 结构体添加speak 方法，输出 xxx是一个好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}

//给 Person 结构体添加jisuan方法，可以计算1+...+1000的结果
func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是=", res)
}

//给 Person 结构体添加jisuan2方法，该方法可以接收一个参数n，计算1+...+n的结果
func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是=", res)
}

//给 Person 结构体添加getSum 方法，可以计算两个数的和，并返回结果
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}




//给 Person 类型绑定一方法
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func main() {
	fmt.Println()

	var p Person
	p.Name = "tom"
	//test方法只能通过 Person 类型的变量来调用，不能直接使用
	p.test()
	//下面的使用方法 都是错误的
	//var dog Dog
	//dog.test()
	//test()
	//调用方法
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	res := p.getSum(10, 20)
	fmt.Println("res=", res)
}
