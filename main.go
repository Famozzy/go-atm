package main

import (
	"fmt"
	"time"

	"go-atm/data"
	"go-atm/handler"
	"go-atm/lib"
	u "go-atm/user"
)

func menu() (int, error) {
	lib.ClearConsole()
	lib.Greeting(u.AuthedUser.GetUsername())
	fmt.Println("+--------------------+")
	fmt.Println("|    GO-ATM          |")
	fmt.Println("+--------------------+")
	fmt.Println("| 1. Show balance    |")
	fmt.Println("| 2. Transfer        |")
	fmt.Println("| 3. Withdraw        |")
	fmt.Println("| 4. Deposit         |")
	fmt.Println("| 5. Logout          |")
	fmt.Println("+--------------------+")
	fmt.Print("Enter the option: ")

	var option int
	_, err := fmt.Scan(&option)

	if err == nil {
		return option, nil
	}

	return 0, err

}

func inputCCNumber() bool {
	attempts := 0
	var ccNumber, pinNumber string

	for {
		if attempts > 0 {
			fmt.Println("info: you have", 3-attempts, "attempts left")
		}

		fmt.Print("Enter your CC Number : ")
		if _, err := fmt.Scan(&ccNumber); err == nil {
			if user, ok := data.Users[ccNumber]; ok {
				fmt.Print("Enter your PIN: ")
				if _, err := fmt.Scan(&pinNumber); err == nil && user.GetPIN() == pinNumber {
					u.AuthedUser = user
					return true
				} else {
					fmt.Println("Invalid PIN")
				}
			} else {
				fmt.Println("Invalid CC Number")
			}
		}

		time.Sleep(2 * time.Second)
		attempts++
		if attempts == 3 {
			return false
		}
	}
}

func main() {
	for ok := inputCCNumber(); ok; {
		lib.ClearConsole()
		if choice, err := menu(); err == nil {
			switch choice {
			case 1:
				handler.ShowBalance()
			case 2:
				handler.Transfer()
			case 3:
				handler.Withdraw()
			case 4:
				handler.Deposit()
			case 5:
				fmt.Println("\nLogout success")
				u.AuthedUser = &u.User{}
				main()
				return
			}
		} else {
			fmt.Println("error:", err.Error())
			time.Sleep(1 * time.Second)
		}
	}

}
