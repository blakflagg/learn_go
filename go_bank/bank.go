package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBallanceFile = "balance.txt"

func main() {
	var choice int
	var accountBal, err = getBalanceFromFile()

	if err != nil {

		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----")
		// panic(err)
	}
	// fmt.Println("Your Choice: ", choice)

	// wantscheckBalance := choice == 1
	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your Balance: %.2f\n ", accountBal)

		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Deposit Amount:")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}

			accountBal += depositAmount
			fmt.Printf("Balance Updated: %.2f\n ", accountBal)
			writeBalanceToFile(accountBal)

		} else if choice == 3 {
			var withdrawlAmount float64
			fmt.Print("Withdrawl Amount:")
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}
			if withdrawlAmount > accountBal {
				fmt.Println("Insufficient Funds")
				continue
			}

			accountBal -= withdrawlAmount
			fmt.Printf("Balance Updated: %.2f\n ", accountBal)
			writeBalanceToFile(accountBal)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing our bank!")
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBallanceFile, []byte(balanceText), 0644)

}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBallanceFile)
	if err != nil {
		return 1000, errors.New("Cannot read from balance file")
	}
	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}

	return balance, nil

}
