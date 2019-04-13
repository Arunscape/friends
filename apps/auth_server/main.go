package main

import (
	"github.com/arunscape/friends/apps/auth_server/database"
	"github.com/arunscape/friends/apps/auth_server/server"
	"github.com/arunscape/friends/commons/server/logger"

	"os"
	"strconv"
)

func main() {
	val := os.Getenv("DID_I_SET_THE_ENVIROMENT_VARIABLES")
	if val != "YES I DID" {
		logger.Error("Enviroment variables not found")
		return
	}

	db := database.NewMySQL()
	db.Open()
	defer db.Close()
	if len(os.Args) > 1 && os.Args[1] == "--setup" {
		db.ResetTheWholeDatabase()
	}

	server.MakeRoutes(db)

	port := 8049
	port_env, _ := strconv.Atoi(os.Getenv("PORT"))
	if port_env != 0 {
		port = port_env
	}
	logger.Info("Starting Server (PORT " + strconv.Itoa(port) + ")")
	server.RunServer(port)
}
