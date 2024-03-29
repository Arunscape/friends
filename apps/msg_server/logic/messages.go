package logic

import (
	"github.com/arunscape/friends/apps/msg_server/database"
	"github.com/arunscape/friends/commons/server/datatypes"
	"github.com/arunscape/friends/commons/server/security"
	"github.com/arunscape/friends/commons/server/web_server"

	"errors"
)

func SendMsgLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*InputSendMsg)

	if !security.ValidateUserToken(data.Tok) {
		return nil, errors.New(web_server.INVALID_SIGNIN_TOKEN)
	}

	// TODO: check if they can send to that group
	if !true {
		return nil, errors.New(web_server.TOKEN_FORBIDDEN)
	}

	db.SendMessage(data.GroupId, data.Msg)
	return nil, nil
}

func GetMsgLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*InputGetMsg)
	msgs := db.QueryMessages(data.GroupId, data.Skip, data.Amount, data.Query)
	return msgs, nil
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
		Msg     datatypes.Message
	}
)
