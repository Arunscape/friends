package database

import (
  "github.com/arunscape/friends/commons/server/datatypes"
)

// AccessObject is the general database access object, all specifc database backends must implement this
type AccessObject interface {
	Open()
	Close()
	ResetTheWholeDatabase()
	GetUserByAuthId(string) (datatypes.User, bool)
	CreateNewUser(datatypes.User)
	CreateNewGroup(datatypes.Group, datatypes.User)
	GetUsersByGroup(datatypes.Group) []datatypes.User
}


