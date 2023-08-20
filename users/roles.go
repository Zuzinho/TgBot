package users

type Role string

const (
	Admin        Role = "Admin"
	User         Role = "User"
	UnwantedUser Role = "Unwanted user"
)

var roleByUser = map[UserName]Role{
	Alena:  User,
	Nikita: User,
}

func GetRole(userName UserName) Role {
	if r, ok := roleByUser[userName]; ok {
		return r
	} else {
		return UnwantedUser
	}
}
