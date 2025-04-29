package main

import (
	"fmt"
	"learn_go/accounts"
)

func main() {
	account := accounts.NewAccount("yejun")
	fmt.Println(account)
}
