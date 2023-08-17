package users

type Role string

const (
	Admin        Role = "Admin"
	User         Role = "User"
	UnwantedUser      = "Unwanted user"
)

var roleByUser = map[string]Role{
	Alena:  User,
	Nikita: User,
}

func GetRole(userName string) Role {
	if r, ok := roleByUser[userName]; ok {
		return r
	} else {
		return UnwantedUser
	}
}
