package handler

import (
	"fmt"
	"go-atm/lib"
	u "go-atm/user"
)

func Withdraw() {
	var amount int
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	u.AuthedUser.Withdraw(amount)
	lib.Next()
}
