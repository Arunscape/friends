package main

import (
	"github.com/arunscape/friends/commons/server/logger"
	"github.com/arunscape/friends/commons/server/mail"

	"encoding/json"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

var (
	SERVER = os.Getenv("EMAIL_SERVER")
	PORT   = os.Getenv("EMAIL_PORT")
	FROM   = os.Getenv("EMAIL_FROM")
	USER   = os.Getenv("EMAIL_USER")
	PASSWD = os.Getenv("EMAIL_PASSWORD")
)

func main() {
	if os.Getenv("DID_I_SET_THE_ENVIROMENT_VARIABLES") != "YES I DID" {
		logger.Error("Enviroment variables not found")
		return
	}
	http.HandleFunc("/send", emailHandler)

	port := 8080
	port_env, _ := strconv.Atoi(os.Getenv("PORT"))
	if port_env != 0 {
		port = port_env
	}
	logger.Info("Starting Server (PORT " + strconv.Itoa(port) + ")")
	logger.Info("Server: "+SERVER, "From: "+FROM, "Port: "+PORT)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func emailHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data mail.EmailMsg
	err := decoder.Decode(&data)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(400)
		w.Write([]byte("400 - Bad Request"))
		return
	}
	err = sendEmail(SERVER, PORT, data.To, FROM, data.Subject, data.Body, USER, PASSWD)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}
	logger.Info("Sent email: " + data.Body)
}

func sendEmail(server, port string, to []string, from, subject, body, auth_usr, auth_passwd string) error {
	return smtp.SendMail(server+":"+port, smtp.PlainAuth("", auth_usr, auth_passwd, server),
		from, to, getEmailBody(to, from, subject, body))
}

func getEmailBody(to []string, from, subject, body string) []byte {
	var sb strings.Builder
	sb.WriteString("To: ")
	sb.WriteString(to[0])
	for i := 1; i < len(to); i++ {
		sb.WriteString(", ")
		sb.WriteString(to[i])
	}
	sb.WriteString("\r\n")

	sb.WriteString("From: ")
	sb.WriteString(from)
	sb.WriteString("\r\n")

	sb.WriteString("Subject: ")
	sb.WriteString(subject)
	sb.WriteString("\r\n")

	sb.WriteString("\r\n")
	sb.WriteString(body)
	sb.WriteString("\r\n")

	return []byte(sb.String())
}
