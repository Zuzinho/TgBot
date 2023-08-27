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

func (mp roleMapType) Add(chatId int64, userName string, role Role) {
	key := roleMapKey{chatId: chatId, userName: userName}

	mp[key] = role
}

func (mp roleMapType) Get(chatId int64, userName string) (Role, bool) {
	key := roleMapKey{chatId: chatId, userName: userName}

	role, ok := mp[key]

	return role, ok
}
