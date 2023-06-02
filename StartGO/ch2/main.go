package main

import (
	"fmt"
	"githubcode/StartGO/ch2/mydict"
)

func main() {
	// account := accounts.NewAccount("nico")
	// account.Deposit(10)
	// err := account.Withdraw(20)
	// if err != nil {
	// 	// log.Fatalln(err)
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	// account.ChangeOwner("John")
	// fmt.Println(account.Balance(), account.Owner())

	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
