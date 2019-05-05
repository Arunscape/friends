package server

import (
	"github.com/arunscape/friends/apps/msg_server/database"
	"github.com/arunscape/friends/commons/server/web_server"

	"net/http"
	"strconv"
)

// MakeRoutes is where all the routing is done. The actual http paths are
// mapped to their handler functions. This should be kept purely for that
// purpose to avoid cluttering it up. That way people will be able to find
// things. The routes are added to DefaultServeMux
func MakeRoutes(db database.AccessObject) {
	http.HandleFunc("/", web_server.NotFoundHandler())

	http.HandleFunc("/msg/new", SendMsgHandler(db))
	http.HandleFunc("/msg/query", GetMsgHandler(db))
}

// RunServer just runs the server on a given port
// it uses go's DefaultServeMux
func RunServer(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
