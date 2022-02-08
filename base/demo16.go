package main

import "fmt"

func main() {
	//for循环
	for i := 1; i<=10; i++ {
		fmt.Println("hellow",i)
	}

	//for循环的第二种写法
	j := 1 //循环变量初始化
	for j <= 10 {
		fmt.Println("hellow",j)
		j++
	}

	//for循环的第三种写法
	k := 1 //循环变量初始化
	for { //这里也等价于 for ; ; {
		if k <= 10 {
			fmt.Println("ok---",k)
		}else{
			break; //break就是跳出这个循环
		}
		k++
	}

	//字符串遍历方式1-传统方式
	var str string = "hello,world!"
	//for i := 0; i <= len(str); i++{
	//	fmt.Printf("%c \n",str[i])
	//}

	//字符串遍历方式1-传统方式
	//var str string = "hello,world!吴广源"
	//str2 := []rune(str) //就是把str转化成 []rune
	//for i := 0; i <= len(str2); i++{
	//	fmt.Printf("%c \n",str2[i])
	//}

	//fmt.Println()
	//字符串遍历方式2  for-range
	str = "abc-ok曹达华"
	for index, val := range str {
		fmt.Printf("index=%d,val=%c \n",index,val)
	}

}