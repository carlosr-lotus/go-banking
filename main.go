package main

import "fmt"

type Account struct {
	id       int
	owner    string
	amount   float64
	password int
}

func (a *Account) Balance() float64 {
	fmt.Print("Current balance: ")
	fmt.Printf("$%.2f", a.amount)

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

	a.amount -= amount
	fmt.Println("Withdraw success!")
	fmt.Printf("Your new balance is: $%.2f\n", a.amount)

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

	var password int
	fmt.Print("Create your account password: ")
	fmt.Scanln(&password)

	acc := Account{1, owner, 0, password}
out:
	for {
		option := BankOptions()

		switch option {
		case 1:
			fmt.Println("************************")
			acc.Balance()
			fmt.Println("\n************************")
		case 2:
			var amount float64
			fmt.Println("************************")
			fmt.Print("Please, insert the amount you wish you to deposit: $")
			fmt.Scanln(&amount)
			acc.Deposit(amount)
			fmt.Println("************************")
		case 3:
			var amount float64
			fmt.Println("************************")
			fmt.Print("Please, insert the amount you wish you to withdraw: $")
			fmt.Scanln(&amount)
			acc.Withdraw(amount)
			fmt.Println("************************")
		case 4:
			fmt.Println("Exiting app...")
			break out
		default:
			fmt.Println("Please, select a valid option.")
		}
	}
}
