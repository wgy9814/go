package main

import "fmt"

func main() {
	//打印1-100之间所有是9的倍数的整数的个数及总和

	//1.使用for循环对max进行遍历
	//2. 当一个数%9 = 0 就是9的倍数
	//3.需要声明两个变量 count和 sum来保存个数和总和
	var max uint64 = 100
	var count uint64 = 0
	var sum uint64 = 0
	var i uint64 = 1
	for ; i <= max; i++ {
		if i % 9 ==0{
			count ++
			sum += i
		}
	}
	fmt.Printf("count=%v,sum=%v \n",count,sum)

	//完成以下的表达式输出，6是可变的
	var n int = 6
	for i := 0; i<=n; i++ {
		fmt.Printf("%v +%v = %v \n", i, n - i, n)
	}

	//使用while方式输出10句 “hello world”
	//循环变量初始化
	var k int = 1
	for {
		if k > 10{//循环条件
			break;//跳出for循环 结束循环
		}
		fmt.Println("hello world",k)
		k++ //循环变量的迭代
	}

	//使用do...while方式输出10句 “hello world”
	var j int = 1
	for {
		fmt.Println("hello world",j)
		j++ //循环变量的迭代
		if j > 10 {//循环条件
			break;//跳出for循环 结束循环
		}
	}

}