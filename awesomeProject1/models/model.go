package models

import "fmt"

type account struct {
	username string
	pwd string
	money int
}

//返回account对象
func NewAccount() *account  {
	return &account{
	}
}

//设置用户名
func (a *account) SetUsername(username string)  {
	if len(username) >=6 && len(username) <=10{
		a.username = username
	}else{
		fmt.Println("用户名必须是6-10位")
	}
}

//getUsername
func (a * account) GetUsername() string {
	return a.username
}

//设置密码
func(a *account) SetPwd(pwd string)  {
	if len(pwd) != 6 {
		fmt.Println("密码必须是6位")
	}else{
		a.pwd = pwd
	}
}

//getPwd
func (a * account) GetPwd() string  {
	return a.pwd
}

//SetMoney
func (a * account) SetMoney(money int)  {
	if money > 20 {
		a.money = money
	}else{
		fmt.Println("余额必须大于20")
	}
}

//getMoney
func (a *account) GetMoney() int  {
	return a.money
}


