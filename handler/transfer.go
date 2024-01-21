package handler

import (
	"fmt"
	"go-atm/data"
	"go-atm/lib"
	u "go-atm/user"
)

func Transfer() {
	var ccNumber string
	var amount int
	fmt.Print("Enter CC Number: ")
	fmt.Scan(&ccNumber)
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)

	if targetUser, ok := data.Users[ccNumber]; ok && targetUser != u.AuthedUser {
		if err := u.AuthedUser.Transfer(targetUser, amount); err != nil {
			fmt.Println("error:", err.Error())
			return
		}
		fmt.Println("Transfer success")
		lib.Next()
	}

}
