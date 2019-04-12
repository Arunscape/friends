package logic

import (
	"github.com/arunscape/friends/apps/api_server/database"

	"encoding/json"
	"errors"
)

type JLogic func(interface{}, database.AccessObject) (interface{}, error)

const INVALID_SIGNIN_TOKEN = "Invalid signin token"
const UNKNOWN = "Something went wrong while processing request"
const INVALID_JSON_RESPONSE = "Failed to encode response as JSON"
const INVALID_JSON_INPUT = "Failed to parse input as JSON"
const USER_DOES_NOT_EXIST = "User does not exist"
const USER_FAILED_TO_CREATE = "Failed to create user"
var reasonStatus = map[string]int{
	INVALID_SIGNIN_TOKEN:  401,
	UNKNOWN:               500,
	INVALID_JSON_RESPONSE: 500,
	INVALID_JSON_INPUT:    400,
	USER_DOES_NOT_EXIST:   401,
  USER_FAILED_TO_CREATE: 500,
}

func JLogicFinalize(msg string) ([]byte, int) {
	status, ok := reasonStatus[msg]
	if !ok {
		msg = UNKNOWN
		status = 500
	}
	return []byte("{\"err\": \"" + msg + "\"}"), status
}

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

func SigninLogic(d interface{}, db database.AccessObject) (interface{}, error) {
	data := d.(*InputSignin)

	gId, isValid := GetGoogleIdFromToken(data.GTok)
	if !isValid {
		return "", errors.New(INVALID_SIGNIN_TOKEN)
	}

	user, found := db.GetUserByAuthId(gId)
	if !found {
		return "", errors.New(USER_DOES_NOT_EXIST)
	}

	val, err := MakeUserFullToken(user)
	return "{\"tok\": \"" + val + "\"}", err
}

func NewUserLogic(d interface{}, db database.AccessObject) (interface{}, error) {
	data := d.(*InputSignup)
	gId, name, email, picture, isValid := GetGoogleInfoFromToken(data.GTok)
  if !isValid {
		return nil, errors.New(USER_FAILED_TO_CREATE)
  }

  user := database.User{
    AuthId: gId,
    Name: name,
    Email: email,
    Picture: picture,
  }
  db.CreateNewUser(user)
	user, found := db.GetUserByAuthId(gId)
  if !found {
		return nil, errors.New(USER_FAILED_TO_CREATE)
  }

	val, err := MakeUserFullToken(user)
	return "{\"tok\": \"" + val + "\"}", err
}

type InputSignup struct {
	GTok string
}

type InputSignin struct {
	GTok string
}
