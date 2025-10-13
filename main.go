package main

import (
	"fmt"
	"password/account"
	"password/files"
)


func main (){
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите Url")
	myAccount, err := account.NewAccountWithTimeStamp(login,password,url)
	if err != nil{
		fmt.Println("Не верный формат URL")
		return
	}
	myAccount.OutputPassword()
	files.WriteFile()
	fmt.Println(myAccount)
}

func promptData(prompt string) string{
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

