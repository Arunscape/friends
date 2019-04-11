package server

import (
	"github.com/arunscape/friends/apps/api_server/database"
	"github.com/arunscape/friends/apps/api_server/logic"

	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
)

func JLogicHandler(fun logic.JLogic, dataType interface{}, db database.AccessObject) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error: ", r)
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
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
