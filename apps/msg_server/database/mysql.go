package database

import (
	"github.com/arunscape/friends/commons/server/datatypes"
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
    gid CHAR(37),
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

func (dao *MySQLAccessObject) SendMessage(gid string, m datatypes.Message) {
	dao.db.Exec("INSERT INTO messages (gid, timestamp, sender, body) VALUES(?, ?, ?, ?)", gid, m.Timestamp, m.Sender, m.Body)
}

func (dao *MySQLAccessObject) QueryMessages(gid string, skip, amount int, text string) []datatypes.Message {
	rows, err := dao.db.Query(`SELECT (timestamp, sender, body) FROM messages
      WHERE gid = ? AND body LIKE '%?%'
      ORDER BY timestamp DESC
      LIMIT ? OFFSET ?`, gid, text, amount, skip)
	if err != nil {
		logger.Error("Failed to query database QueryMessages: ", err)
		return make([]datatypes.Message, 0)
	}
	msgs := make([]datatypes.Message, 0)
	for rows.Next() {
		var m datatypes.Message
		err = rows.Scan(&m.Timestamp, &m.Sender, &m.Body)
		if err == nil {
			msgs = append(msgs, m)
		}
	}
	return msgs
}
