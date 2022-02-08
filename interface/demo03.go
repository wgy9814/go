package main

import "fmt"

//interface 类型默认是一个指针（引用类型）
type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}

//如果需要实现 AInterface，就需要将 BInterface CInterface的方法都实现.

type Stu struct {

}
func (stu Stu) test01() {
	fmt.Println("手机开始工作.....")
}

func (stu Stu) test02() {
	fmt.Println("手机开始工作.....")
}
func (stu Stu) test03() {
	fmt.Println("手机开始工作.....")
}

type T interface {

}

type Usb interface {
	Say()
}

type Stus struct {

}
func (this *Stus) Say() {
	fmt.Println("Say()")
}

func main() {
	//测试
	//先创建结构体变量
	var stu Stu
	var a AInterface = stu
	a.test01()

	//空接口interface{}没有任何方法，所有类型都实现了空接口，即我们可以如何把一个变量赋给空接口
	var t T = stu
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2,t)


	var stus Stus = Stus{}
	//var u Usb = stus//错误！ 会报 Stu类型没有实现Usb接口
	var u Usb = &stus
	u.Say()
	fmt.Println("hero",u)
}
