package server

import (
	//"github.com/arunscape/friends/commons/server/logger"
	"github.com/arunscape/friends/commons/server/web_server"
	"github.com/arunscape/friends/apps/msg_server/database"
	"github.com/arunscape/friends/apps/msg_server/logic"

	"net/http"
)


func GetMsgHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return web_server.JLogicHandler(logic.GetMsgLogic, &logic.InputGetMsg{}, db)
}

func SendMsgHandler(db database.AccessObject) func(http.ResponseWriter, *http.Request) {
	return web_server.JLogicHandler(logic.SendMsgLogic, &logic.InputSendMsg{}, db)
}
