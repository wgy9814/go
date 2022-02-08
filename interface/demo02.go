package main

import "fmt"

//定义一个接口
type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type BInterface interface {
	Hello()
}

type Monster struct {

}

func (m Monster) Hello() {
	fmt.Println("Monster Hello")
}
func (m Monster) Say() {
	fmt.Println("Monster Say")
}

func main() {
	//接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
	var stu Stu //结构体变量，实现了 Say() 实现了 AInterface
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say() //integer Say i =10

	//Monster 实现了 AInterface 和 BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}
