package logic

import (
	//"github.com/arunscape/friends/commons/server/web_server"
	"github.com/arunscape/friends/apps/msg_server/database"

	//"encoding/json"
	//"errors"
)

func SendMsgLogic(d interface{}, db_dat interface{}) (interface{}, error) {
    // db := db_dat.(database.AccessObject)
	//data := d.(*InputSendMsg)
    return nil, nil // TODO: make nil return a 204 in the JLogicHttpWrapper
}

func GetMsgLogic(d interface{}, db_dat interface{}) (interface{}, error) {
     //db := db_dat.(database.AccessObject)
	//data := d.(*InputGetMsg)
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
