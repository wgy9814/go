package main

import "fmt"

func main() {

	//有一个数列：白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王
	//猜数游戏：从键盘中任意输入一个名称，判断数列中是否包含此名称【顺序查找】
	//要求: 如果找到了，就提示找到，并给出下标值

	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名")
	fmt.Scanln(&heroName)

	//顺序查找：第一种方式
	//for i := 0; i < len(names); i++ {
	//	if heroName == names[i]{
	//		fmt.Printf("找到%v,下标为%v \n", heroName, i)
	//		break
	//	} else if i ==  len(names) - 1 {
	//		fmt.Printf("没有找到%v \n", heroName)
	//	}
	//}

	//顺序查找：第2种方式
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i]{
			index = i //将找到的值对应下标 赋给 index
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v,下标为%v \n", heroName, index)
	} else {
		fmt.Printf("没有找到%v \n", heroName)
	}
}
