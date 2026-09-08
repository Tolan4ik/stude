package main

import "fmt"

type Account struct {
	ID      int
	Balance int
}

type PremiumAccount struct {
	Account
	CashbackPercent int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func main() {
	p := PremiumAccount{}
	p.Deposit(100)
	fmt.Println("Новый баланс:", p.Balance)
}
