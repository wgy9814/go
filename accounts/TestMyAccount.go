package main

import "fmt"

//基本数据转成string使用
func main() {
	// 声明一个变量，保存接收用户输入的选项
	key := ""
	// 声明一个变量，控制是否退出循环
	loop := true

	// 定义账户余额
	balance := 10000.0
	//// 每次收支的金额
	money := 0.0
	//// 每次收支的说明
	note := ""
	//// 定义一个变量，记录是否有收支行为
	flag := false
	// 收支的详情，使用字符串来记录
	// 当有收支时，只需要对details进行拼接处理即可
	details := "收支\t账户金额\t收支金额\t说明"

	// 显示这个主菜单
	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("          1 收支明细")
		fmt.Println("          2 登记收入")
		fmt.Println("          3 登记支出")
		fmt.Println("          4 退出")
		fmt.Println("          请选择(1-4):")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("----------当前收支明细记录----------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细。。。。来一笔")
			}

		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money //修改账户余额
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			// 将这个收入情况，拼接到detalis变量
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额的金额不足")
				break
			}
			balance	-= money
			fmt.Println("本次支入说明:")
			fmt.Scanln(&note)
			// 将这个收入情况，拼接到detalis变量
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗？ y/n")
			choice := ""
			for  {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入  y/n")
			}
			if choice == "y" {
				loop = false
			}
		default :
			fmt.Println("请输入正确的选项(1-4)")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件的使用......")

}