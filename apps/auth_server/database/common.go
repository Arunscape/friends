package database

import (
	"github.com/arunscape/friends/apps/auth_server/logger"

	"os/exec"
  "strings"
)

// AccessObject is the general database access object, all specifc database backends must implement this
type AccessObject interface {
	Open()
	Close()
	ResetTheWholeDatabase()
	GetUserByAuthId(string) (User, bool)
	CreateNewUser(User)
}

// User is the defintion of that data that a user should contain
type User struct {
	Name    string
	Email   string
	Picture string
	Id      string
	AuthId  string
}

// UUID should really be made more general, but this was so easy
// It generates a 128bit unique identifier, stored as a string 
// https://en.wikipedia.org/wiki/Universally_unique_identifier
func UUID() string {
	// Works on linux only, probably. Sorry
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		logger.Error("Failed to create uuid")
	}
	return strings.Trim(string(out), " \t\n\f\r\v") // strip all whitespace
}
