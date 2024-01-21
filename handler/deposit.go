package handler

import (
	"fmt"
	"go-atm/lib"
	u "go-atm/user"
)

func Deposit() {
	var amount int
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	u.AuthedUser.Deposit(amount)
	lib.Next()
}
