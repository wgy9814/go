package model

import "fmt"

//定义一个结构体
type person struct {
	Name string
	age int
	sal float64
}


//写一个工厂模式来解决 类似构造函数

func NewPerson(name string) *person {
	return &person{
		Name : name,
	}
}

//为了访问age 和 sal 我们编写一对setxx的方法和getxx方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确...")
		//程序员给一个默认值
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal < 40000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确...")
		//程序员给一个默认值
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}


