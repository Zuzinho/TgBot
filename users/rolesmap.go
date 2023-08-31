package users

type roleMapKey struct {
	chatId   int64
	userName string
}

type roleMapType map[roleMapKey]Role

var roleMap roleMapType

func init() {
	roleMap = make(roleMapType)
}

func AddRole(chatId int64, userName string, role Role) {
	key := roleMapKey{chatId: chatId, userName: userName}

	roleMap[key] = role
}

func GetRole(chatId int64, userName string) (Role, bool) {
	key := roleMapKey{chatId: chatId, userName: userName}

	role, ok := roleMap[key]

	return role, ok
}
