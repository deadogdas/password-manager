package main

import (
	"fmt"
	"password/account"
	"password/files"
)


func main (){
	createAccount()
}

func createAccount(){
	// files.ReadFile()
	// files.WriteFile("Привет!", "file.txt")
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите Url")
	myAccount, err := account.NewAccount(login,password,url)
	if err != nil{
		fmt.Println("Не верный формат URL")
		return
	}
	file, err := myAccount.ToBytes()
		if err != nil{
		fmt.Println("Не удалось приобразовать в JSON")
		return
	} 
	files.WriteFile(file,"data.json")
}

func promptData(prompt string) string{
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

