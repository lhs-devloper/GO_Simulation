package main

import (
	"fmt"
	"log"

	"github.com/lhs/learngo/accounts"
)

//private = 소문자
//public = 대문자
func main() {
	account := accounts.NewAccount("lhs")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.String())
}
