package users

import "fmt"

type UnknownRoleError struct {
	error
	roleInt  int
	roleRole Role
}

func (err UnknownRoleError) String() string {
	return fmt.Sprintf("unknown user role %d or %s", err.roleInt, err.roleRole)
}

func NewUnknownRoleError(role Role) UnknownRoleError {
	return UnknownRoleError{
		roleRole: role,
	}
}
