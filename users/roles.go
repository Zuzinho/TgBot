package users

type Role string

const (
	Admin        Role = "Admin"
	User         Role = "User"
	UnwantedUser      = "Unwanted user"
)

var roleByUser = map[string]Role{
	Alena:  User,
	Nikita: Admin,
}

func GetRole(userName string) Role {
	return User
	if r, ok := roleByUser[userName]; ok {
		return r
	} else {
		return UnwantedUser
	}
}
