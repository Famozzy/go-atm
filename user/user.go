package structures

import (
	"fmt"
	"go-atm/lib"
)

var AuthedUser *User

type IUser interface {
	Transfer(user *User, amount int)
	Withdraw(amount int)
	Deposit(amount int)
}

type User struct {
	username string
	pin      string
	balance  int
	cash     int
}

func NewUser(username, pin string, balance, cash int) *User {
	return &User{username, pin, balance, cash}
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetPIN() string {
	return u.pin
}

func (u User) GetCash() int {
	return u.cash
}

func (u User) GetBalance() int {
	return u.balance
}

func (u *User) Transfer(user *User, amount int) error {
	if amount > u.balance {
		return fmt.Errorf("insufficient balance")
	}

	user.balance += amount
	u.balance -= amount

	username := lib.CensorName(user.username)
	fmt.Println("Transfer success to", username, "with amount", amount)
	return nil
}

func (u *User) Withdraw(amount int) error {
	if amount > u.balance {
		return fmt.Errorf("insufficient balance")
	}

	u.balance -= amount
	u.cash += amount
	fmt.Println("Withdraw success with amount", amount)
	return nil
}

func (u *User) Deposit(amount int) error {
	if amount > u.cash {
		return fmt.Errorf("insufficient cash")
	}

	u.balance += amount
	u.cash -= amount
	fmt.Println("Deposit success with amount", amount)
	return nil
}
