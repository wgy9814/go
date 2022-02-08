package main

import "fmt"

//二分查找的函数
/*
二分查找的思路：比如我们要查找的数是 findval
1. arr是一个有序数组，并且是从小到大排列
2.先找到中间下标 middle = (leftIndex + rightIndex) / 2 然后让中间下标的值和findval进行比较
2.1.如果 arr[middle] > findval,就应该向 leftIndex ---- （middle - 1）
2.2.如果 arr[middle] < findval,就应该向 rightIndex ---- （middle + 1）
2.3.如果 arr[middle] == findval，就找到了
2.4 上面的2.1 2.2 2.3会递归执行
3 什么情况下 就说明找不到（退出递归条件）
if leftIndex > rightIndex{
	找不到
	return
}
 */
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	//判断leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		//说明我们要找的数在  应该在leftIndex ---- （middle - 1）
		BinaryFind(arr, leftIndex, middle - 1, findVal)
	} else if  (*arr)[middle] < findVal {
		//说明我们要找的数在  应该在rightIndex ---- （middle + 1）
		BinaryFind(arr,  middle + 1, rightIndex, findVal)
	} else {
		//找到了
		fmt.Printf("找到了，下标为%v \n", middle)
	}


}

func main() {

	//二分查找的函数
	arr := [6]int{1, 8, 10, 89, 1000, 1234}

	BinaryFind(&arr, 0, len(arr) - 1, 89)
}
