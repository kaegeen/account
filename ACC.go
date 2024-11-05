package main

import "fmt"

// BankAccount struct to represent a bank account
type BankAccount struct {
	balance float64
}

// Method to deposit an amount into the account
func (account *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		account.balance += amount
		fmt.Printf("Successfully deposited $%.2f\n", amount)
	} else {
		fmt.Println("Deposit amount must be positive.")
	}
}

// Method to withdraw an amount from the account
func (account *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= account.balance {
		account.balance -= amount
		fmt.Printf("Successfully withdrew $%.2f\n", amount)
	} else if amount > account.balance {
		fmt.Println("Insufficient balance.")
	} else {
		fmt.Println("Withdrawal amount must be positive.")
	}
}

// Method to display the current balance
func (account *BankAccount) DisplayBalance() {
	fmt.Printf("Current balance: $%.2f\n", account.balance)
}

func main() {
	var choice int
	var amount float64

	// Create a new bank account
	account := BankAccount{balance: 0}

	for {
		fmt.Println("\nBank Account Menu:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Display Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.DisplayBalance()
		case 4:
			fmt.Println("Exiting. Thank you!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
