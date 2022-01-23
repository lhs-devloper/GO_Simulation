package accounts

import (
	"errors"
	"fmt"
)

//Account Struct(구조체 연습)
//owner = 사람, balance: 금액
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

//New Account creates Account(새로 생성)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//receiver
//규칙: struct 첫 글자를 따서 소문자로 지어야함
//Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance of your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOnwer of the Account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Onwer of the Account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas:", a.Balance())
}
