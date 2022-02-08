package main

import "fmt"

func main() {
	//统计3个班的成绩情况 每个班有5个同学
	//求出每个班的平均分和所有班级的总分（成绩从键盘输入）

	//分析实现思路
	//1.统计一个班成绩情况。每个班有5个同学，求出该办的平均分【成绩从键盘输入】 =》先易后难
	//2. 学生数就是5个【先死后活】
	//3.声明一个sum 统计班级的总分

	//分析实现思路2
	//1.统计3个班成绩情况。每个班有5个同学，求出该办的平均分【成绩从键盘输入】
	//2.j代表几个班
	//3.定义一个变量存放总成绩

	//分析实现思路3
	//把代码做活
	//定义两个变量 表示班级的个数和班级的人数

	//统计3个班的及格人数，每个班有5名同学
	//分析实现思路
	//声明passcount 用于保存及格人数

	var classNum int = 2
	var stuNum int = 5
	var totalSum float64 = 0.0
	var passcount int = 0
	for j := 1; j <=classNum; j++ {
		sum := 0.0
		for i := 1; i <= stuNum; i++ {
			var scroe float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩 \n", j, i)
			fmt.Scanln(&scroe)
			//累计总分
			sum += scroe
			//判断是否及格
			if scroe >= 60 {
				passcount ++
			}
		}
		fmt.Printf("第%d个班的平均分是%v \n", j, sum / float64(stuNum))
		//将每个班的成绩累加到totalSum
		totalSum += sum
	}
	fmt.Printf("每个班级的总成绩%v 所有班级的平均分是%v \n", totalSum, totalSum / (float64(stuNum * classNum)))
	fmt.Printf("及格人数为%v \n", passcount)

}