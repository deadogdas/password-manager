package main

import (
	"fmt"
	"password/account"

	"github.com/fatih/color"
)


func main (){
	vault := account.NewVault()
	fmt.Println("__Менеджер паролей__")
Menu:
	for{
		variant := getMenu()
		switch variant{
		case 1 :
			createAccount(vault)
		case 2:
			fiendAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break Menu
		}
	}

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


func promptData(prompt string) string{
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func fiendAccount(vault *account.Vault){
	url := promptData("Введите URL для поиска")
	accounts := vault.FiendAccountsByUrl(url)
	if len(accounts) == 0{
		color.Red("Аккаунтов не найдено")
	}
	for _, acaccount := range accounts{
		acaccount.Output()
	}
}

func deleteAccount(vault *account.Vault){
	url := promptData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted{
		color.Green("Удаленно")
	}else{
		color.Red("Не найдено")
	}


}

func createAccount(vault *account.Vault){
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите Url")
	myAccount, err := account.NewAccount(login,password,url)
	if err != nil{
		color.Red("Не верный формат URL")
		return
	}
	vault.AddAccount(*myAccount)
}