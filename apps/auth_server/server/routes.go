package server

import (
	"github.com/arunscape/friends/apps/auth_server/database"

	"net/http"
	"strconv"
)

// MakeRoutes is where all the routing is done. The actual http paths are
// mapped to their handler functions. This should be kept purely for that
// purpose to avoid cluttering it up. That way people will be able to find
// things. The routes are added to DefaultServeMux
func MakeRoutes(db database.AccessObject) {
	http.HandleFunc("/", NotFoundHandler())

	http.HandleFunc("/user/signin", ValidateHandler(db))

	http.HandleFunc("/test/signin", GoogleWebSigninHandler())
}

// RunServer just runs the server on a given port
// it uses go's DefaultServeMux
func RunServer(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
