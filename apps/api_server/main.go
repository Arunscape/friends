package main

import (
	"github.com/arunscape/friends/apps/api_server/database"
	"github.com/arunscape/friends/apps/api_server/server"

	"fmt"
	"os"
)

func main() {
	val := os.Getenv("DID_I_SET_THE_ENVIROMENT_VARIABLES")
	if val != "YES I DID" {
		fmt.Println("Enviroment variables not found")
		return
	}

	db := database.NewMySQL()
	db.Open()
	defer db.Close()
	if len(os.Args) > 1 && os.Args[1] == "--setup" {
		db.ResetTheWholeDatabase()
	}

	fmt.Println("### Starting Server ###")
	server.MakeRoutes(db)
	server.RunServer(8049)
}
