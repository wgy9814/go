package main

import "fmt"

type Person struct {
	Name string
}



//函数
//对于普通函数，接受者为值类型时，不能将指针类型的数据直接传递，反之亦然
func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

//对于方法（如struct的方法）
//接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样可以
func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03 = ", p.Name)
}

//地址拷贝
func (p *Person) test04() {
	p.Name = "mary"
	fmt.Println("test04 = ", p.Name)
}


func main() {
	p := Person{"tom"}
	test01(p)
	test02(&p)

	p.test03()
	fmt.Println("main() p.name=", p.Name)

	(&p).test03() //从形式上是传入地址，但是本质还是值拷贝
	fmt.Println("main() p.name=", p.Name)

	(&p).test04() //
	fmt.Println("main() p.name=", p.Name)
	p.test04()
}
