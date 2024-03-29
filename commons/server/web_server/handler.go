package web_server

import (
	"github.com/arunscape/friends/commons/server/logger"
	"io/ioutil"
	"net/http"
	"runtime/debug"
)

// JLogicHandler is for creating our logic handlers. It catches panics, and
// reads the http.Request to extract relivent information, and writes the
// returned json to the response object. Basically the common code for each
// handler
func JLogicHandler(fun JLogic, dataType interface{}, db interface{}) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if Cors(w, req) {
			return
		}
		logger.Info(req.Method, " ", req.URL)
		defer func() {
			if r := recover(); r != nil {
				logger.Error("Error:", r, "\n", "stacktrace from panic: \n"+string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(errorResponse("Something went horribly wrong"))
			}
		}()

		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorResponse("Json not well formatted for this request"))
		}

		jsonOut, status := JLogicHttpWrapper(fun, dataType, body, db)
		w.WriteHeader(status)
		w.Write(jsonOut)
	}
}

func errorResponse(in string) []byte {
	return []byte("{\"err\": \"" + in + "\"}")
}

// NotFoundHandler is called when an endpoint that doesn't exist is accessed
func NotFoundHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if Cors(w, req) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(errorResponse("Method not found"))
	}
}

func Cors(w http.ResponseWriter, req *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*req).Method == "OPTIONS" {
		return true
	}
	return false
}

type NoData struct{}
