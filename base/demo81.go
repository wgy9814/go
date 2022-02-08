package main

import (
	"fmt"
)


type A struct {
	Name string
	Age int
}

func (a *A) Sayok() {
	fmt.Println("A Sayok", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) Sayok() {
	fmt.Println("B Sayok", b.Name)
}




func main() {

	//结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或小写的字段，方法，都可以使用
	//var b B
	//b.A.Name = "tom"
	//b.A.Age = 20
	//b.A.Sayok()
	//b.A.hello()

	//上面的写法还可以简化
	//编译器会先看b对应的类型有没有Name，如果有，则直接调用B类型的Name字段
	//如果没有就去看B中嵌入的匿名结构体A有没有声明 Name 字段 如果找不到 报错
	//b.Name = "smith"
	//b.Age = 20
	//b.Sayok()
	//b.hello()

	//相同字段或方法 才去就近原则访问
	var b B
	b.Name = "jack"
	b.A.Name = "smith"
	b.Age = 100
	b.Sayok()
	b.A.Sayok()
	b.hello()
}
