package database

import (
	"github.com/arunscape/friends/apps/logger"

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
	dao.db.Exec("DROP TABLE IF EXISTS permissions")
	dao.db.Exec("DROP TABLE IF EXISTS groups")
	dao.db.Exec("DROP TABLE IF EXISTS users_groups")
	dao.db.Exec(`CREATE TABLE groups(
    id CHAR(37),
    name VARCHAR(256),
    PRIMARY KEY(id))`)

	dao.db.Exec(`CREATE TABLE users_groups(
    uid CHAR(37),
    gid CHAR(37))`)

	dao.db.Exec(`CREATE TABLE permissions(
    uid CHAR(37),
    name VARCHAR(127))`)

	dao.db.Exec(`CREATE TABLE users(
    id CHAR(37),
    authId VARCHAR(256),
    name VARCHAR(256),
    email VARCHAR(256),
    picture VARCHAR(512),
    PRIMARY KEY(id))`)

	dao.CreateNewUser(User{AuthId: "49", Name: "Testy McTestface", Email: "testy@test.test", Picture: "https://i.guim.co.uk/img/media/ddda0e5745cba9e3248f0e27b3946f14c4d5bc04/108_0_7200_4320/master/7200.jpg?width=620&quality=45&auto=format&fit=max&dpr=2&s=dff8678a6e1cdd5716fe6c49767bac9a"})
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
func (dao *MySQLAccessObject) CreateNewUser(user User) {
	_, isUser := dao.GetUserByAuthId(user.AuthId)
	if !isUser {
		dao.db.Exec("INSERT INTO users(id, name, email, picture, authId) VALUES(?, ?, ?, ?, ?)", UUID(), user.Name, user.Email, user.Picture, user.AuthId)
	}
}

func (dao *MySQLAccessObject) GetUserByAuthId(id string) (User, bool) {
	var user User
	found := true
	err := dao.db.QueryRow("select id, name, email, picture, authId from users where authId = ?", id).Scan(&user.Id, &user.Name, &user.Email, &user.Picture, &user.AuthId)
	if err != nil {
		logger.Debug("Failed to query database GetUserByAuthId: ", err)
		found = false
	}
	dao.getGroupsByUser(&user)
    dao.getPermissionsByUser(&user)
	return user, found
}

func (dao *MySQLAccessObject) CreateNewGroup(g Group, u User) {
	gid := UUID()
	dao.db.Exec("INSERT INTO groups(id, name) VALUES(?, ?)", gid, g.Name)
	dao.db.Exec("INSERT INTO users_groups(uid, gid) VALUES(?, ?)", u.Id, gid)
}

func (dao *MySQLAccessObject) GetUsersByGroup(g Group) []User {
	rows, err := dao.db.Query(`
      SELECT u.id, u.name, u.email, u.picture, u.authId FROM groups g
      JOIN users_groups ug ON g.id = ug.gid
      JOIN users u ON ug.uid = u.id
      WHERE g.id = ?`, g.Id)
	if err != nil {
		logger.Error("Failed to query database GetUsersByGroup: ", err)
		return make([]User, 0)
	}
	users := make([]User, 0)
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Picture, &u.AuthId)
		if err == nil {
			users = append(users, u)
		}
	}

	return users
}

func (dao *MySQLAccessObject) getGroupsByUser(u *User) {
	rows, err := dao.db.Query(`
      SELECT g.id, g.name FROM groups g
      JOIN users_groups ug ON g.id = ug.gid
      JOIN users u ON ug.uid = u.id
      WHERE u.id = ?`, u.Id)
	if err != nil {
		logger.Error("Failed to query database getGroupsByUser:", err)
		return
	}
	groups := make([]Group, 0)
	for rows.Next() {
		var g Group
		err = rows.Scan(&g.Id, &g.Name)
		if err == nil {
			groups = append(groups, g)
		}
	}
	u.Groups = groups
}

func (dao *MySQLAccessObject) getPermissionsByUser(u *User) {
	rows, err := dao.db.Query(`
      SELECT name FROM permissions
      WHERE uid = ?`, u.Id)
	if err != nil {
		logger.Error("Failed to query database getPermissionsByUser:", err)
		return
	}
	permissions := make([]string, 0)
	for rows.Next() {
		var p string
		err = rows.Scan(&p)
		if err == nil {
			permissions  = append(permissions, p)
		}
	}
	u.Permissions = permissions
}
