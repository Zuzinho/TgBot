package users

import (
	"ZuzinhoBot/database/selector"
	"log"
)

type Role string

const (
	Admin        Role = "Admin"
	User         Role = "User"
	UnwantedUser Role = "Unwanted user"
)

func NewRole(chatId int64, userName string) (role Role, err error) {
	role, exists := roleMap.Get(chatId, userName)
	if exists {
		log.Printf("Got from map role '%s'", role)

		return role, nil
	}

	id, err := selector.SelectUserRole(chatId, userName)
	if err != nil {
		return
	}

	switch id {
	case 1:
		role = UnwantedUser
	case 2:
		role = User
	case 3:
		role = Admin
	default:
		return "", NewUnknownRoleErrorByInt(id)
	}

	roleMap.Add(chatId, userName, role)

	return
}
