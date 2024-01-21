package data

import (
	u "go-atm/user"
)

var Users = map[string]*u.User{
	"123654": u.NewUser("Faidil", "1234", 100000, 50000),
	"123098": u.NewUser("Famozzy", "4321", 140000, 100000),
	"123443": u.NewUser("Savid", "1111", 200000, 150000),
	"123321": u.NewUser("John", "2222", 300000, 200000),
}
