package main

import "fmt"

type Account struct {
	id    int
	owner string
}

func (a *Account) Deposit(amount float64) bool {
	fmt.Println(amount)
	fmt.Println("Deposit successful!")

	return true
}

func main() {
	fmt.Println("******************")
	fmt.Println("    GO BANKING    ")
	fmt.Println("******************")

	newAccount := Account{1, "Lotus"}

	newAccount.Deposit(120)
}
