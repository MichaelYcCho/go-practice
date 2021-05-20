package main

import (
	"fmt"

	"github.com/michael_cho77/go-practice/go-scrapper/accounts"
)

func main() {
	account := accounts.NewAccount("michael")
	account.Deposit(10)

	fmt.Println(account.String())
}
