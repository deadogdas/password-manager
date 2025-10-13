package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)

var letterRunes =[]rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-*!")


type Account struct{
	login string
	password string
	url string
}

type AccountWithTimeStamp struct{
	createdAt time.Time
	updateAt time.Time
	Account
}

func (acc *Account) OutputPassword(){
	color.Cyan(acc.login)
}

func (acc *Account) generatePassword(n int){
	res:= make([]rune,n)
	for i := range res{
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp,error){
	if login == ""{
		return nil, errors.New("INVALID_LOGIN")
	}
	_,err := url.ParseRequestURI(urlString)
	if err !=nil{
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updateAt: time.Now(),
		Account: Account{
			url: urlString,
			login: login,
			password: password,
		},
	}
	if password == ""{
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

