package logic

import (
	"github.com/arunscape/friends/apps/msg_server/database"

	"encoding/json"
	"errors"
)

type JLogic func(interface{}, database.AccessObject) (interface{}, error)

const UNKNOWN = "Something went wrong while processing request"

var reasonStatus = map[string]int{
	UNKNOWN: 500,
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

func JLogicHttpWrapper(fun JLogic, in interface{}, data []byte, db database.AccessObject) ([]byte, int) {
	err := json.Unmarshal(data, in)
	if err != nil {
		return JLogicFinalize(INVALID_JSON_INPUT)
	}

	val, err := fun(in, db)
	if err != nil {
		return JLogicFinalize(err.Error())
	}

	bytes, err := json.Marshal(val)
	if err != nil {
		return JLogicFinalize(INVALID_JSON_RESPONSE)
	}
	return bytes, 200
}

func SendMsgLogic(d interface{}, db database.AccessObject) (interface{}, error) {
	data := d.(*InputSendMsg)
    return nil, nil // TODO: make nil return a 204 in the JLogicHttpWrapper
}

func GetMsgLogic(d interface{}, db database.AccessObject) (interface{}, error) {
	data := d.(*InputGetMsg)
    return nil, nil // TODO: make nil return a 204 in the JLogicHttpWrapper
}

type (
	InputGetMsg struct {
		Tok     string
		GroupId string
		Query   string
		Skip    int
		Amount  int
	}

	InputSendMsg struct {
		Tok     string
		GroupId string
		Msg     database.Message
	}
)
