package main

import (
	"fmt"
	"password/account"

)


func main (){
	fmt.Println("__Менеджер паролей__")
Menu:
	for{
		variant := getMenu()
		switch variant{
		case 1 :
			createAccount()
		case 2:
			fiendAccount()
		case 3:
			deleteAccount()
		case 4:
			break Menu
		}
	}
	createAccount()
}

func getMenu() int{
	var variant int 
	fmt.Println("Выберите вариант:")
	fmt.Println("1.Создать аккаунт")
	fmt.Println("2.Найти аккаунт")
	fmt.Println("3.Удалить аккаунт")
	fmt.Println("4.Выход")
	fmt.Scan(&variant)
	return variant
}

func createAccount(){
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите Url")
	myAccount, err := account.NewAccount(login,password,url)
	if err != nil{
		fmt.Println("Не верный формат URL")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string{
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func fiendAccount(){

}

func deleteAccount(){

}