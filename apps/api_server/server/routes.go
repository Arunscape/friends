package server

import (
	"github.com/arunscape/friends/apps/api_server/database"

	"net/http"
	"strconv"
)

func MakeRoutes(db database.AccessObject) {
	http.HandleFunc("/", NotFoundHandler())

	http.HandleFunc("/user/new", NewUserHandler(db))
	http.HandleFunc("/user/signin", SigninHandler(db))
}

func RunServer(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
