package main

import (
	"fmt"
)


/*
定义一个结构体 Account
 */
type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

//方法1 存款
func (account *Account) SaveMoney(monney float64, pwd string) {
	//看下输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看下存款金额是否正确
	if monney <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += monney
	fmt.Print("存款成功--")
}

//方法2 取款
func (account *Account) WithDraw(monney float64, pwd string) {
	//看下输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看下存款金额是否正确
	if monney <= 0 || monney > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance -= monney
	fmt.Print("取款成功--")
}

//方法3 查询余额
func (account *Account) Query(pwd string) {
	//看下输入密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为=%v 余额=%v \n", account.AccountNo, account.Balance)
}



func main() {
	//测试一把
	account := Account{
		AccountNo: "g132",
		Pwd: "123456",
		Balance: 500.0,
	}
	account.Query("123456")
	account.SaveMoney(200.0, "123456")
	account.Query("123456")

	account.WithDraw(350.0, "123456")
	account.Query("123456")
}
