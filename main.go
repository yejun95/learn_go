package main

import (
	"fmt"
	"learn_go/accounts"
)

func main() {
	account := accounts.NewAccount("yejun")
	account.Deposit(10)
	fmt.Println(account)
	fmt.Println(account.Balance())
}
