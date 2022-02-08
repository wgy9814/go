package main

import "fmt"

//定义一个cat结构体，将cat的各个字段属性信息，放入到cat结构体进行管理 值类型
type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
	Scores [3]int //字段是数组
}

func main() {
	//1.使用变量或者数组来解决养猫的问题,不利于数据的管理和维护。因为名字,年龄,颜色都是
	//属于一只猫,但是这里是分开保存。
	//2.如果我们希望对一只猫的属性(名字、年龄,颜色)进行操作(绑定方法), 也不好处理。
	//3.引出我们要讲解的技术-》结构体。

	//Golang 也支持面向对象编程(OOP),但是和传统的面向对象编程有区别,并不是纯粹的面向对象语言。所以我们说 Golang 支持面向对象编程特性是比较准确的。
	//Golang 没有类(class),Go 语言的结构体(struct)和其它编程语言的类(class)有同等的地位,你可以理解 Golang 是基于 struct 来实现 OOP 特性的。
	//Golang 面向对象编程非常简洁,去掉了传统 OOP 语言的继承、方法重载、构造函数和析构函数、隐藏的 this 指针等等。
	//Golang 仍然有面向对象编程的继承,封装和多态的特性,只是实现的方式和其它 OOP 语言不一样,比如继承 :Golang 没有 extends 关键字,继承是通过匿名字段来实现。
	//Golang 面向对象(OOP)很优雅,OOP 本身就是语言类型系统(type system)的一部分,通过接口(interface)关联,耦合性低,也非常灵活。
	//后面同学们会充分体会到这个特点。也就是说在 Golang 中面向接口编程是非常重要的特性。

	// 创建一个cat的变量
	var cat1 Cat
	fmt.Printf("cat1的地址=%p\n", &cat1)
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃饭<..>>>>"
	fmt.Println("cat1=", cat1)

	fmt.Println("猫猫的信息如下")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("hobby=", cat1.Hobby)


}
