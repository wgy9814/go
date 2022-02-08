package model

import (
	"fmt"
)


/*
定义一个结构体 Account
*/
type account struct {
	accountNo string
	pwd string
	balance float64
}

//写一个工厂模式 类似构造函数

func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对...")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码的长度不对...")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数目不对...")
		return nil
	}
	return &account{
		accountNo : accountNo,
		pwd : pwd,
		balance : balance,
	}
}




//方法1 存款
func (account *account) SaveMoney(monney float64, pwd string) {
	//看下输入密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看下存款金额是否正确
	if monney <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.balance += monney
	fmt.Print("存款成功--")
}

//方法2 取款
func (account *account) WithDraw(monney float64, pwd string) {
	//看下输入密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看下存款金额是否正确
	if monney <= 0 || monney > account.balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.balance -= monney
	fmt.Print("取款成功--")
}

//方法3 查询余额
func (account *account) Query(pwd string) {
	//看下输入密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为=%v 余额=%v \n", account.accountNo, account.balance)
}



