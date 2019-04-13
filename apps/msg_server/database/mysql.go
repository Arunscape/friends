package database

import (
	"github.com/arunscape/friends/commons/server/logger"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"os"
)

type MySQLAccessObject struct {
	db *sql.DB
}

func NewMySQL() *MySQLAccessObject {
	return &MySQLAccessObject{}
}

func (dao *MySQLAccessObject) ResetTheWholeDatabase() {
	dao.db.Exec("DROP TABLE IF EXISTS messages")

	dao.db.Exec(`CREATE TABLE messages(
    timestamp BIGINT,
    sender CHAR(37),
    body VARCHAR(256))`)

	logger.Info("Database reset and ready to go")
}

func (dao *MySQLAccessObject) Open() {
	dataString := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWD") + "@tcp(" + os.Getenv("DB_LOC") + ")/" + os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", dataString)
	if err != nil {
		logger.Error("Failed to connect to database: ", dataString, err)
	}

	dao.db = db
}

func (dao *MySQLAccessObject) Close() {
	dao.db.Close()
}
