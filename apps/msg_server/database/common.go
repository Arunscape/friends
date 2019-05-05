package database

import (
	"github.com/arunscape/friends/commons/server/datatypes"
)

// AccessObject is the general database access object, all specifc database backends must implement this
type AccessObject interface {
	Open()
	Close()
	ResetTheWholeDatabase()
	SendMessage(string, datatypes.Message)
	QueryMessages(string, int, int, string) []datatypes.Message
}
