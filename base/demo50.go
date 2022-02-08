package main

import "fmt"

func BubbleSort(arr *[5]int){
	fmt.Println("排序前的arr=", (*arr))
	temp := 0 //临时变量
	for i := 0;i < len(*arr) - 1; i++ {
		for j := 0;j < len(*arr) - 1 - i; j++ {
			if (*arr)[j] > (*arr)[j+1]{
				//交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}
	}
}

func main() {

	/*
	1.排序的介绍
	排序是将一群数据，依指定的顺序进行排列的过程。

	排序的分类：
	（1）内部排序
	将需要处理的所有数据都加载到内部存储器中进行排序。
	包括（交换式排序，选择式排序和插入式排序）

	（2）外部排序
	数据量过大，无法全部加载到内存中，需要借助外部存储进行排序。
	包括（合并排序法和直接合并排序法）
	 */

	//冒泡排序

	arr := [5]int{24,69,80,57,13}
	BubbleSort(&arr);
	fmt.Println("排序后的arr=", arr)
}
