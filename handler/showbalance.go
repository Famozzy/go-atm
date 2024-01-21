package handler

import (
	"fmt"
	"go-atm/lib"
	u "go-atm/user"
)

func ShowBalance() {
	fmt.Println("Your balance is", u.AuthedUser.GetBalance())
	fmt.Println("Your cash is", u.AuthedUser.GetCash())
	lib.Next()
}
