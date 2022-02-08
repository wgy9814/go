package main

import "fmt"


type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int //指针
	slice []int //切片
	map1 map[string]string//map
}



func main() {
	//方式1
	var p1 Person
	fmt.Println("p1=", p1)

	//方式2
	p2 := Person{}
	p2.Name = "吴广源"
	p2.Age = 18
	fmt.Println(p2)
	//p2 := Person{"赤峰黄金和",20}

	//方式3-&
	var p3 *Person = new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = 'smith' 也可以写成 p3.Name = "smith"
	//原因：go的设计者 为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理
	//会给 p3 加上 取值运算 (*p3).Name = 'smith'
	(*p3).Name = "smith"
	p3.Name = "john"

	(*p3).Age = 53
	p3.Age = 100
	fmt.Println(*p3)

	//方式4-{}
	//案例：
	var person *Person = &Person{}
	//var person *Person = &Person{"mary", 60} //也可以直接赋值
	//因为 person 是一个指针，因此标准的给字段赋值方式
	//(*person).Name = "scott" 也可以写成 person.Name = "scott"
	//原因和上面一样 	会给 person 加上 取值运算 (*person).Name = 'scott'
	(*person).Name = "scott"
	person.Name = "scott--"
	(*person).Age = 18
	person.Age = 77
	fmt.Println(*person)




}
