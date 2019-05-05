package server

import (
	"github.com/arunscape/friends/apps/auth_server/database"
	"github.com/arunscape/friends/apps/auth_server/logic"
	"github.com/arunscape/friends/commons/server/logger"
	"github.com/arunscape/friends/commons/server/web_server"

	"net/http"
	"strings"
)

// ValidateHandler creates a handler for signing in users using the standard jLogic handler pattern
// This is the handler bound to route /validate/BlahBlahBlahTokenGoesHereBlahBlahBlab
// it expects the following json object in the request body
// {
//    "Tok": "aaa.bbb.ccc"
// }
// Returns
// <p>Thanks!</p>
func ValidateHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info(req.Method, " ", req.URL)

		secret := strings.Split(req.URL.String(), "/")[2]
		logger.Debug("Upgrading token with secret: ", secret)
		ok := db.UpgradeToken(secret)

		res.Header().Set("Content-Type", "application/json")
		if ok {
			res.WriteHeader(200)
			res.Write([]byte("Success, you can now login regularly"))
		} else {
			res.WriteHeader(404)
			res.Write([]byte("Sorry, there was a problem. Did the link expire?"))
		}
	}
}

// SigninHandler: /signin
// Expects:
// {
//    "Email": "user@example.com"
// }
// Returns:
// {
//    "Tok": "aaa.bbb.ccc"
// }
func SigninHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return web_server.JLogicHandler(logic.SigninLogic, &logic.Email{}, db)
}

// SignupHandler: /signup
// Expects:
// {
//    "Email": "user@example.com"
//    "Name":  "Testy McTestface"
//    "Picture": "example.com/picture_url.png"
// }
// Returns:
// {
//    "Tok": "aaa.bbb.ccc"
// }
func SignupHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return web_server.JLogicHandler(logic.SignupLogic, &logic.Signup{}, db)
}

// IsUserHandler: /isuser
// Expects:
// {
//    "Email": "user@example.com"
// }
// Returns: isUserReal ? 204 : 404
func IsUserHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return web_server.JLogicHandler(logic.IsUserLogic, &logic.Email{}, db)
}

// SignoutHandler: /signout
// Expects:
// {
//    "Tok": "aaa.bbb.ccc"
// }
// Returns: isTokenValid ? 204 : 403
func SignoutHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return web_server.JLogicHandler(logic.SignoutLogic, &logic.Token{}, db)
}

// UpgradeHandler: /upgrade
// Expects:
// {
//    "Tok": "aaa.bbb.ccc"
// }
// Returns:
// {
//    "Tok": "aaa.bbb.ccc"
// }
func UpgradeHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return web_server.JLogicHandler(logic.UpgradeLogic, &logic.Token{}, db)
}
