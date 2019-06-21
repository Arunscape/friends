package mail

import (
	"github.com/arunscape/friends/commons/server/logger"

	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type EmailMsg struct {
	To      []string
	Subject string
	Body    string
}

func SendEmail(to []string, subject string, body string) bool {
	mail := EmailMsg{To: to, Subject: subject, Body: body}
	b, err := json.Marshal(&mail)
	if err != nil {
		return false
	}
	resp, err := http.Post("http://email."+os.Getenv("DOMAIN")+"/send", "application/json", bytes.NewBuffer(b))
	logger.Debug("Sent email: ", resp, err)
	return err == nil && resp.StatusCode < 300 && resp.StatusCode >= 200
}
