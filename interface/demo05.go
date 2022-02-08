package main

import (
	"fmt"
)

// Monkey 结构体
type Monkey struct {
	Name string
}

//type BirdAble interface {
//	Flying()
//}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

//LittleMonkey 结构体
type LittleMonkey struct {
	Monkey //继承
}

//让LittleMonkey 实现 BirdAble

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习 ，会飞")
}
func main() {
	//创建一个 LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey {
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
}
