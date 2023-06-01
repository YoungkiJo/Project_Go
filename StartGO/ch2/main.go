package main

import (
	"fmt"
	accounts "githubcode/StartGO/ch2/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	account.ChangeOwner("John")
	fmt.Println(account.Balance(), account.Owner())
}
