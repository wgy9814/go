package main

import "fmt"

func main() {
	//使用for 循环完成下面的案例，可以接收一个整数，表示层数，打印金字塔

	//编程思路
	//1.打印一个矩形
	/*
		***
		***
		***
	*/
	//i表示层数
	for i := 1; i <= 3; i ++ {
		//j表示每层打印多少*
		for j := 1; j <= 3; j ++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//2.打印半个金字塔
	/*
		*
		**
		***
	 */

	for i := 1; i <= 3; i ++ {
		//j表示每层打印多少*
		for j := 1; j <= i; j ++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//3.打印一个金字塔
	/*
		  *  规律：2 * 层数 - 1  空格 规律 总层数-当前层数
		 ***
		*****
	 */
	//4.将层数做成一个变量
	//var totalLevel int
	//var totalLevel int = 9
	//for i := 1; i <= totalLevel; i ++ {
	//	//先打印空格
	//	for k := 1; k <= totalLevel - i; k ++ {
	//		fmt.Print(" ")
	//	}
	//	//j表示每层打印多少*
	//	for j := 1; j <= 2 * i - 1; j ++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	//5.打印一个空心金字塔
	/*
		  *  规律：2 * 层数 - 1  空格 规律 总层数-当前层数
		 * *
		*****
		分析：在我们给每一行打印*号时，需要考虑的是打印 * 还是打印空格
		我们的分析结果是，每层的第一个和最后一个是打印* , 其他打印输出空格
	    最后一行全打印*
	*/
	//4.将层数做成一个变量
	//var totalLevel int

	var totalLevel int = 9
	for i := 1; i <= totalLevel; i ++ {
		//先打印空格
		for k := 1; k <= totalLevel - i; k ++ {
			fmt.Print(" ")
		}
		//j表示每层打印多少*
		for j := 1; j <= 2 * i - 1; j ++ {
			if j == 1 || j == 2 * i - 1 || i == totalLevel {
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//打印九九乘法表
	//i表示层数
	var num int = 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j * i)
		}
		fmt.Println()
	}
}