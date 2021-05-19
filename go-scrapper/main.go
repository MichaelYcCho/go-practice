package main

import (
	"fmt"

	"github.com/michael_cho77/go-practice/go-scrapper/accounts"
)

func main() {
	account := accounts.NewAccount("michael")
	fmt.Println(account)

}
