package main

import (
	"fmt"
	"learn_go/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("yejun")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(account.String())
}
