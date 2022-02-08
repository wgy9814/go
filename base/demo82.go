package main

import (
	"fmt"
)


type A struct {
	Name string
	Age int
}

type B struct {
	Name string
	Score float64
}


type C struct {
	A
	B
	//Name string
}

type D struct {
	a A //有名结构体
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {//多重继承
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age int
}

type E struct {
	Monster
	int //匿名字段基本数据类型 只能有一个
	n int
}



func main() {

	var c C
	//如果c 没有 Name 字段，而A和B有 Name，这时就必须指定匿名结构体名字来区分
	//所以 c.Name 会报错，这个规则对于方法也是一样
	c.A.Name = "tom"
	fmt.Printf("C")

	//如果D 中是一个有名结构体，则访问有名结构体的字段时，就必须带上有名结构体的名字
	//比如 d.a.Name
	var d D
	d.a.Name = "jack"

	//嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
	tv := TV{ Goods{"电视剧001", 5000.9}, Brand{"海尔", "山东"}, }
	tv2 := TV{
		Goods{
			Price: 5000.9,
			Name : "电视剧002",
		},
		Brand{
			Name: "夏普",
			Address: "北京",
		},
	}
	fmt.Println("tv=",tv)
	fmt.Println("tv2=",tv2)

	tv3 := TV2{ &Goods{"电视剧003", 5000.9}, &Brand{"小米", "miui"}, }
	tv4 := TV2{
		&Goods{
			Name : "电视剧004",
			Price: 9000.9,
		},
		&Brand{
			Name: "长虹",
			Address: "四川",
		},
	}
	fmt.Println("tv3=", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4=", *tv4.Goods, *tv4.Brand)

	//演示一下匿名字段时基本数据类型的使用
	var e E
	e.Name = "火精灵"
	e.Age = 20
	e.int = 20
	e.n = 40
	fmt.Println("e=", e)
}
