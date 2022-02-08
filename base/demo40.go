package main

import "fmt"

func main() {
	//定义一个数组
	var hens [7]float64
	//给数组每个元素赋值
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	hens[6] = 150.0

	//遍历数组求出总体重
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	//求出平均体重
	avgWeight := fmt.Sprintf("%.2f", totalWeight / float64(len(hens)))
	fmt.Printf("totalWeight=%v avgWeight=%v", totalWeight, avgWeight)

	var intArr [3]int //int占4个字节
	//当我们定义完数组后，其实数组的各个元素都有默认值0
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)

	//1.数组的地址可以通过&intArr来获取
	//2.数组的第一个元素的地址，就是数组的首地址
	//3.数组的各个元素的地址间隔是依据数组类型决定的，比如int64->8 int32->4
	fmt.Printf("intArr的地址=%p intArr[0]的地址=%p intArr[1]的地址=%p intArr[1]的地址=%p \n",
		&intArr, &intArr[0],  &intArr[1],  &intArr[2])

	//从终端循环输入5个成绩，保存到float64数组，并输出
	//var scroe [5]float64
	//for i := 0; i < len(scroe); i++ {
	//	fmt.Printf("请输入第%d个元素的值\n", i+1)
	//	fmt.Scanln(&scroe[i])
	//}
	//
	////变量数组打印
	//for i := 0; i < len(scroe); i++ {
	//	fmt.Printf("scroe[%d]=%v\n", i, scroe[i])
	//}

	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=",numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02=",numArr02)
	//这里[...]是规定的写法
	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03=",numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04=",numArr04)

	//类型推导
	numArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("numArr05=",numArr05)


}
