package database

import (
	"github.com/arunscape/friends/apps/auth_server/logger"

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
	dao.db.Exec("DROP TABLE IF EXISTS users")
	dao.db.Exec(`CREATE TABLE users(
    id CHAR(37),
    authId VARCHAR(256),
    name VARCHAR(256),
    email VARCHAR(256),
    picture VARCHAR(512),
    PRIMARY KEY(id)
  )`)
  dao.CreateNewUser(User{AuthId: "49", Name: "Testy McTestface", Email: "testy@test.test", Picture: "https://i.guim.co.uk/img/media/ddda0e5745cba9e3248f0e27b3946f14c4d5bc04/108_0_7200_4320/master/7200.jpg?width=620&quality=45&auto=format&fit=max&dpr=2&s=dff8678a6e1cdd5716fe6c49767bac9a"})
	logger.Info("Database reset and ready to go")
}

func (dao *MySQLAccessObject) Open() {
	dataString := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWD") + "@tcp(" + os.Getenv("DB_LOC") + ")/" + os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", dataString)
	if err != nil {
		logger.Error("Failed to connect to database: ", dataString)
	}

	dao.db = db
}

func (dao *MySQLAccessObject) Close() {
	dao.db.Close()
}
func (dao *MySQLAccessObject) CreateNewUser(user User) {
	dao.db.Exec("INSERT INTO users(id, name, email, picture, authId) VALUES(?, ?, ?, ?, ?)", UUID(), user.Name, user.Email, user.Picture, user.AuthId)
}

func (dao *MySQLAccessObject) GetUserByAuthId(id string) (User, bool) {
	var user User
  found := false
	rows, err := dao.db.Query("select id, name from users where authId = ?", id)
	if err != nil {
		logger.Error("Failed to query database")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
      logger.Error("ERROR", err)
		}
    found = true
	}
	err = rows.Err()
	if err != nil {
    logger.Error("ERROR", err)
	}
	return user, found
}
