package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func (a *Account) CheckPassword(password int) bool {
	if password == a.password {
		return true
	} else {
		return false
	}
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

func ReturnIcons(icon string, quantity int) string {
	var formattedIcon string
	for c := 0; c < quantity; c++ {
		formattedIcon += icon
	}
	return formattedIcon
}

func main() {
	fmt.Println(ReturnIcons("*", 18))
	fmt.Println("    GO BANKING    ")
	fmt.Println(ReturnIcons("*", 18))

	reader := bufio.NewReader(os.Stdin)

	var owner string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&owner)

	var password int
	for {
		fmt.Print("Create your account password: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error when creating your password. Please, try again.")
			continue
		}

		input = strings.TrimSpace(input)

		if len(input) != 4 {
			fmt.Println("Your password must be 4 numbers only. Please, try again.")
		}

		isValid := true
		for _, char := range input {
			if !('0' <= char && char <= '9') {
				isValid = false
				break
			}
		}

		if isValid {
			intValue, err := strconv.Atoi(input)

			if err == nil {
				password = intValue
				fmt.Printf("Your password is: %d\n", intValue)
				break
			}
		}

		fmt.Println("Error when creating password. Please, try again.")
	}

	acc := Account{1, owner, 0, password}
out:
	for {
		option := BankOptions()

		switch option {
		case 1:
			fmt.Println(ReturnIcons("*", 24))
			acc.Balance()
			fmt.Print("\n")
			fmt.Println(ReturnIcons("*", 24))
		case 2:
			var amount float64
			fmt.Println(ReturnIcons("*", 24))
			fmt.Print("Please, insert the amount you wish you to deposit: $")
			fmt.Scanln(&amount)
			acc.Deposit(amount)
			fmt.Println(ReturnIcons("*", 24))
		case 3:
			var amount float64
			var password int
			fmt.Println(ReturnIcons("*", 24))
			fmt.Printf("To continue, insert your password: ")
			fmt.Scanln(&password)
			check := acc.CheckPassword(password)

			if check {
				fmt.Print("Please, insert the amount you wish you to withdraw: $")
				fmt.Scanln(&amount)
				acc.Withdraw(amount)
				fmt.Println(ReturnIcons("*", 24))
			} else {
				fmt.Println("Process interrupted. Wrong password.")
				fmt.Println(ReturnIcons("*", 24))
			}
		case 4:
			fmt.Println("Exiting app...")
			break out
		default:
			fmt.Println("Please, select a valid option.")
		}
	}
}
