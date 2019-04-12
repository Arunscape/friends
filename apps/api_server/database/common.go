package database

import (
	"github.com/arunscape/friends/apps/api_server/logger"

	"os/exec"
)

type AccessObject interface {
	Open()
	Close()
	ResetTheWholeDatabase()
	GetUserByAuthId(string) (User, bool)
	CreateNewUser(User)
}

type User struct {
	Name    string
	Email   string
	Picture string
	Id      string
	AuthId  string
}

func UUID() string {
	// Works on linux only, probably. Sorry
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		logger.Error("Failed to create uuid")
	}
	return string(out)
}
