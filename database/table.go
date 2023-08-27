package database

type TableName string

const (
	PhrasesTableName        TableName = "phrases"
	BoringStoriesTableName  TableName = "boring_stories"
	FunnyStoriesTableName   TableName = "funny_stories"
	DataTypesTableName      TableName = "data_types"
	RolesTableName          TableName = "roles"
	UsersTableName          TableName = "users"
	UsersStatisticTableName TableName = "users_statistic"
)

func (t TableName) IsValid() bool {
	return (t == PhrasesTableName) || (t == BoringStoriesTableName) ||
		(t == FunnyStoriesTableName) || (t == DataTypesTableName) ||
		(t == RolesTableName) || (t == UsersTableName) || (t == UsersStatisticTableName)
}

func (t TableName) IsAccessedForUser() bool {
	return (t == PhrasesTableName) || (t == BoringStoriesTableName) || (t == FunnyStoriesTableName)
}
