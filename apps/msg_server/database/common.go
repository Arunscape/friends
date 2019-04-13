package database

import (
	"github.com/arunscape/friends/commons/server/logger"

	"os/exec"
	"strings"
    "time"
)

// AccessObject is the general database access object, all specifc database backends must implement this
type AccessObject interface {
	Open()
	Close()
	ResetTheWholeDatabase()
}

type Message struct {
  Body string
  Sender string
  Timestamp uint64
}

func (m *Message) SetTimestamp() {
  m.Timestamp = uint64(time.Now().UnixNano() / int64(time.Millisecond))
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
