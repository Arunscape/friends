package server

import (
	"github.com/arunscape/friends/apps/api_server/database"
	"github.com/arunscape/friends/apps/api_server/logic"
	"github.com/arunscape/friends/apps/api_server/logger"

  "os"
	"io/ioutil"
	"net/http"
	"runtime/debug"
)

func JLogicHandler(fun logic.JLogic, dataType interface{}, db database.AccessObject) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
    logger.Info(req.Method, " ", req.URL)
		defer func() {
			if r := recover(); r != nil {
				logger.Error("Error:", r, "\n", "stacktrace from panic: \n" + string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(ErrorResponse("Something went horribly wrong"))
			}
		}()

		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ErrorResponse("Json not well formatted for this request"))
		}

    logger.Debug("Request body: ", string(body))
		jsonOut, status := logic.JLogicHttpWrapper(fun, dataType, body, db)
		w.WriteHeader(status)
		w.Write(jsonOut)
	}
}

func ErrorResponse(in string) []byte {
	return []byte("{\"err\": \"" + in + "\"}")
}

func NotFoundHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(ErrorResponse("Method not found"))
	}
}
func SigninHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return JLogicHandler(logic.SigninLogic, &logic.InputSignin{}, db)
}
func NewUserHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return JLogicHandler(logic.NewUserLogic, &logic.InputSignup{}, db)
}

func GoogleWebSigninHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html lang=\"en\"> <head> <meta name=\"google-signin-scope\" content=\"profile email\"> <meta name=\"google-signin-client_id\" content=\"" + 
    os.Getenv("GOOGLE_CLOUD_ID") + 
    "\"> <script src=\"https://apis.google.com/js/platform.js\" async defer></script> </head> <body> <div class=\"g-signin2\" data-onsuccess=\"onSignIn\" data-theme=\"dark\"></div><script>function onSignIn(googleUser){var profile=googleUser.getBasicProfile(); console.log(\"ID: \" + profile.getId()); console.log('Full Name: ' + profile.getName()); console.log('Given Name: ' + profile.getGivenName()); console.log('Family Name: ' + profile.getFamilyName()); console.log(\"Image URL: \" + profile.getImageUrl()); console.log(\"Email: \" + profile.getEmail()); var id_token=googleUser.getAuthResponse().id_token; console.log(\"ID Token: \" + id_token);}</script> </body></html>"))
  }
}


