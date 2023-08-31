package users

import (
	"ZuzinhoBot/database/selector"
)

type Role string

const (
	Admin        Role = "Admin"
	User         Role = "User"
	UnwantedUser Role = "Unwanted user"
)

func NewRole(chatId int64, userName string) (role Role, err error) {
	role, ok := GetRole(chatId, userName)
	if ok {
		return
	}

	roleBD, err := selector.SelectUserRole(chatId, userName)
	if err != nil {
		return
	}

	role = Role(roleBD)

	AddRole(chatId, userName, role)

	return
}
