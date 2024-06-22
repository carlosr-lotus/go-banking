package main

import "fmt"

type Account struct {
	id       int
	owner    string
	amount   float64
	password int
}

func (a *Account) Balance() float64 {
	fmt.Println("************************")
	fmt.Print("Current balance: ")
	fmt.Printf("$ %.2f", a.amount)
	fmt.Println("\n************************")

	return a.amount
}

func (a *Account) Deposit(amount float64) bool {
	a.amount += amount
	fmt.Println(amount)
	fmt.Println("Deposit successful!")

	return true
}

func (a *Account) Withdraw(amount float64) bool {
	if amount > a.amount {
		fmt.Println("Withdraw error! Balance not enough for withdraw operation.")
		return false
	}

	fmt.Println("Withdraw success!")
	return true
}

func BankOptions() int {
	var option int
	fmt.Println("\nChoose option below: ")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit funds")
	fmt.Println("3. Withdraw funds")
	fmt.Println("4. Exit")

	fmt.Scanln(&option)

	return option
}

func main() {
	fmt.Println("******************")
	fmt.Println("    GO BANKING    ")
	fmt.Println("******************")

	var owner string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&owner)

	acc := Account{1, "Lotus", 0, 1234}

out:
	for {
		option := BankOptions()

		switch option {
		case 1:
			acc.Balance()
		case 2:
			acc.Deposit(120)
		case 3:
			acc.Withdraw(150)
		case 4:
			fmt.Println("Exiting app...")
			break out
		default:
			fmt.Println("Please, select a valid option.")
		}
	}
}
