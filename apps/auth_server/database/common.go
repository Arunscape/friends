package database

import (
	"github.com/arunscape/friends/commons/server/datatypes"
)

// AccessObject is the general database access object, all specifc database backends must implement this
type AccessObject interface {
	Open()
	Close()
	ResetTheWholeDatabase()
	GetUserByEmail(string) (datatypes.User, bool)
	CreateNewUser(*datatypes.User)
	CreateNewGroup(datatypes.Group, datatypes.User)
	GetUsersByGroup(datatypes.Group) []datatypes.User
	AddUserValidation(*datatypes.User, string)
	SignInUser(*datatypes.User)
	SignOutUser(*datatypes.User)
	UpgradeToken(string) bool
	SaveUserSettings(datatypes.User) bool
}
