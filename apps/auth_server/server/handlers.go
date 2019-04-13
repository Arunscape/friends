package server

import (
	"github.com/arunscape/friends/commons/server/web_server"
	"github.com/arunscape/friends/apps/auth_server/database"
	"github.com/arunscape/friends/apps/auth_server/logic"

	"net/http"
	"os"
)

// ValidateHandler creates a handler for signing in users using the standard jLogic handler pattern
// This is the handler bound to route /users/signin
// it expects the following json object in the request body
// {
//    "GTok": "aaa.bbb.ccc"
// }
func ValidateHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return web_server.JLogicHandler(logic.ValidateUserLogic, &logic.InputSignin{}, db)
}

// GoogleWebSigninHandler creates the handler for testing the signin button
// This is the handler bound to the route /test/signin
func GoogleWebSigninHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html lang=\"en\"> <head> <meta name=\"google-signin-scope\" content=\"profile email\"> <meta name=\"google-signin-client_id\" content=\"" +
			os.Getenv("GOOGLE_CLOUD_ID") +
			"\"> <script src=\"https://apis.google.com/js/platform.js\" async defer></script> </head> <body> <div class=\"g-signin2\" data-onsuccess=\"onSignIn\" data-theme=\"dark\"></div><script>function onSignIn(googleUser){var profile=googleUser.getBasicProfile(); console.log(\"ID: \" + profile.getId()); console.log('Full Name: ' + profile.getName()); console.log('Given Name: ' + profile.getGivenName()); console.log('Family Name: ' + profile.getFamilyName()); console.log(\"Image URL: \" + profile.getImageUrl()); console.log(\"Email: \" + profile.getEmail()); var id_token=googleUser.getAuthResponse().id_token; console.log(\"ID Token: \" + id_token);}</script> </body></html>"))
	}
}
