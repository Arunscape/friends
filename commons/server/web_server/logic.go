package web_server

import (
	"github.com/arunscape/friends/commons/server/logger"

	"encoding/json"
)

type JLogic func(interface{}, interface{}) (interface{}, error)

const UNKNOWN = "Something went wrong while processing request"
const INVALID_SIGNIN_TOKEN = "Invalid signin token"
const INVALID_JSON_RESPONSE = "Failed to encode response as JSON"
const INVALID_JSON_INPUT = "Failed to parse input as JSON"
const USER_DOES_NOT_EXIST = "User does not exist"
const USER_NOT_FOUND = "User was not found"
const USER_FAILED_TO_CREATE = "Failed to create user"
const TOKEN_FORBIDDEN = "You are not authorized to do that"

var reasonStatus = map[string]int{
	UNKNOWN:               500,
	INVALID_SIGNIN_TOKEN:  401,
	INVALID_JSON_RESPONSE: 500,
	INVALID_JSON_INPUT:    400,
	USER_DOES_NOT_EXIST:   401,
	USER_NOT_FOUND:        404,
	USER_FAILED_TO_CREATE: 500,
	TOKEN_FORBIDDEN:       403,
}

// JLogicFinalize matches returned error messages to the proper status codes,
// this allows unique error messages without having to worry about passing
// multiple values everywhere
func JLogicFinalize(msg string) ([]byte, int) {
	status, ok := reasonStatus[msg]
	if !ok {
		msg = UNKNOWN
		status = 500
	}
	return []byte("{\"err\": \"" + msg + "\"}"), status
}

// JLogicHttpWrapper is a convinence tool, It deals with the parsing the JSON
// input, as well as encoding the returned struct to json. This means our logic
// functions can just deal with structs instead of json. Unfortanatly, it's not
// typesafe yet. If I figure that out, I'll be pretty excited. It returns the
// json out, and the http status code

func JLogicHttpWrapper(fun JLogic, in interface{}, data []byte, db interface{}) ([]byte, int) {
	logger.Debug("Request body: ", string(data), " | ", in)
	err := json.Unmarshal(data, in)
	if err != nil {
		return JLogicFinalize(INVALID_JSON_INPUT)
	}
	val, err := fun(in, db)
	if err != nil {
		return JLogicFinalize(err.Error())
	}

	if val == nil {
		return nil, 204
	}

	bytes, err := json.Marshal(val)
	if err != nil {
		return JLogicFinalize(INVALID_JSON_RESPONSE)
	}
	return bytes, 200
}
