package database

import (
	"fmt"
	"os/exec"
)

type AccessObject interface {
	Open()
	Close()
	ResetTheWholeDatabase()
	GetUserByAuthId(string) User
	CreateNewUser(User)
}

type User struct {
	Name   string
	Id     string
	AuthId string
}

func UUID() string {
	// Works on linux only, probably. Sorry
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		fmt.Println("Failed to create uuid")
	}
	return string(out)
}
