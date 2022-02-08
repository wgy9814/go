package main

import "fmt"


/*

 */
type Stu struct {
	Name string
	Age int
}



func main() {
	//方式1
	//在创建结构体变量时，就直接指定变量的值
	var stu1 = Stu{"小明", 20}
	stu2 := Stu{"小明-", 30}

	//在创建结构体变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
	var stu3 = Stu{
		Name : "电视剧",
		Age : 20,
	}
	stu4 := Stu{
		Age : 99,
		Name : "发多少",
	}
	fmt.Println(stu1, stu2, stu3, stu4)

	//方式2,返回结构体的指针类型（！！！）
	var stu5 = &Stu{"小望", 65} //stu5-->地址 ---》 结构体数据[xx,xx]
	stu6 := &Stu{"小望", 69}

	//在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
	var stu7 = &Stu{
		Name : "小李",
		Age : 20,
	}
	stu8 := &Stu{
		Age : 99,
		Name : "小李",
	}

	fmt.Println(*stu5, *stu6, *stu7, *stu8)



}
