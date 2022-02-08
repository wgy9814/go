package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//我们为了生成一个随机数，还需要rand设置一个种子
	//time.Now().Unix() : 返回一个从1970:01:01的0时0分0秒到现在的秒数
	//rand.Seed(time.Now().Unix())
	////如何随机生成1-100整数
	//n := rand.Intn(100) + 1 //[0 100]
	//fmt.Println(n)

	//随机生成1-100的一个数，直到生成99这个数，看看你一共用了几次
	//分析思路
	//无限循环的控制，不停生成随机数，当生成99时,break退出这个无限循环
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		//如何随机生成1-100整数
		n := rand.Intn(100) + 1 //[0 100]
		fmt.Println("n=",n)
		count++
		if n == 99{
			break;//退出这个无限循环
		}
	}
	fmt.Println("生成99一共用了 ", count)

	//这里演示一下指定标签的形式来使用 break
	label2: //设置一个标签
	for i := 0; i < 4; i++ {
		//label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break //break 默认会跳出最近的for循环
				break label2
			}
			fmt.Println("j=",j)
		}
	}

	//continue案例
	//label2: //设置一个标签
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j=",j)
		}
	}

}