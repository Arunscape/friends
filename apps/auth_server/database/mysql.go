package database

import (
	"github.com/arunscape/friends/commons/server/datatypes"
	"github.com/arunscape/friends/commons/server/logger"
	"github.com/arunscape/friends/commons/server/utils"

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
	dao.db.Exec("DROP TABLE IF EXISTS groups")
	dao.db.Exec("DROP TABLE IF EXISTS permissions")
	dao.db.Exec("DROP TABLE IF EXISTS users_groups")
	dao.db.Exec("DROP TABLE IF EXISTS validations")
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
    name VARCHAR(256),
    email VARCHAR(256),
    picture VARCHAR(512),
    PRIMARY KEY(id))`)

	dao.db.Exec(`CREATE TABLE validations(
    uid CHAR(37),
    isValidated BOOL,
    isSignedIn BOOL,
    secret CHAR(37),
    PRIMARY KEY(uid))`)

	dao.CreateNewUser(&datatypes.User{Name: "Testy McTestface", Email: "testy@test.test", Picture: "https://i.guim.co.uk/img/media/ddda0e5745cba9e3248f0e27b3946f14c4d5bc04/108_0_7200_4320/master/7200.jpg?width=620&quality=45&auto=format&fit=max&dpr=2&s=dff8678a6e1cdd5716fe6c49767bac9a"})
	logger.Info("Database reset and ready to go")
}

func (dao *MySQLAccessObject) Open() {
	dataString := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWD") + "@tcp(" + os.Getenv("DB_LOC") + ")/" + os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", dataString)
	logger.Info("Connecting to database: ", dataString)
	if err != nil {
		logger.Error("Failed to connect to database: ", dataString, err)
	}
	dao.db = db
}

func (dao *MySQLAccessObject) Close() {
	dao.db.Close()
}

func (dao *MySQLAccessObject) CreateNewUser(user *datatypes.User) {
	_, isUser := dao.GetUserByEmail(user.Email)
	if !isUser {
		dao.db.Exec("INSERT INTO users(id, name, email, picture) VALUES(?, ?, ?, ?)", utils.UUID(), user.Name, user.Email, user.Picture)
		logger.Debug("Successfully inserted new user: ", user)
	} else {
		logger.Debug("Failed to insert new user: ", user)
	}
}

func (dao *MySQLAccessObject) GetUserByEmail(id string) (datatypes.User, bool) {
	var user datatypes.User
	found := true
	err := dao.db.QueryRow("select id, name, email, picture from users where email = ?", id).Scan(&user.Id, &user.Name, &user.Email, &user.Picture)
	if err != nil {
		logger.Debug("Failed to query database GetUserByEmail with email (", id, "): ", err)
		found = false
	}
	dao.getGroupsByUser(&user)
	dao.getPermissionsByUser(&user)
	dao.getValidationsByUser(&user)
	return user, found
}

func (dao *MySQLAccessObject) CreateNewGroup(g datatypes.Group, u datatypes.User) {
	gid := utils.UUID()
	dao.db.Exec("INSERT INTO groups(id, name) VALUES(?, ?)", gid, g.Name)
	dao.db.Exec("INSERT INTO users_groups(uid, gid) VALUES(?, ?)", u.Id, gid)
}

func (dao *MySQLAccessObject) GetUsersByGroup(g datatypes.Group) []datatypes.User {
	rows, err := dao.db.Query(`
      SELECT u.id, u.name, u.email, u.picture, u.authId FROM groups g
      JOIN users_groups ug ON g.id = ug.gid
      JOIN users u ON ug.uid = u.id
      WHERE g.id = ?`, g.Id)
	if err != nil {
		logger.Error("Failed to query database GetUsersByGroup: ", err)
		return make([]datatypes.User, 0)
	}
	users := make([]datatypes.User, 0)
	for rows.Next() {
		var u datatypes.User
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Picture, &u.AuthId)
		if err == nil {
			users = append(users, u)
		}
	}

	return users
}

func (dao *MySQLAccessObject) getGroupsByUser(u *datatypes.User) {
	rows, err := dao.db.Query(`
      SELECT g.id, g.name FROM groups g
      JOIN users_groups ug ON g.id = ug.gid
      JOIN users u ON ug.uid = u.id
      WHERE u.id = ?`, u.Id)
	if err != nil {
		logger.Error("Failed to query database getGroupsByUser:", err)
		return
	}
	groups := make([]datatypes.Group, 0)
	for rows.Next() {
		var g datatypes.Group
		err = rows.Scan(&g.Id, &g.Name)
		if err == nil {
			groups = append(groups, g)
		}
	}
	u.Groups = groups
}

func (dao *MySQLAccessObject) getPermissionsByUser(u *datatypes.User) {
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
			permissions = append(permissions, p)
		}
	}
	u.Permissions = permissions
}
func (dao *MySQLAccessObject) getValidationsByUser(usr *datatypes.User) {
	dao.db.QueryRow("select secret, isSignedIn, isValidated from validations where uid = ?", usr.Id).Scan(
		&usr.Secret, &usr.IsSignedIn, &usr.IsValidated)
}
func (dao *MySQLAccessObject) setValidationsByUser(usr *datatypes.User) {
	dao.db.Exec("REPLACE INTO validations(uid, isSignedIn, isValidated, secret) VALUES(?, ?, ?, ?)",
		usr.Id, usr.IsSignedIn, usr.IsValidated, usr.Secret)
}
func (dao *MySQLAccessObject) AddUserValidation(usr *datatypes.User, secret string) {
	usr.IsSignedIn = false
	usr.IsValidated = true
	usr.Secret = secret
	dao.setValidationsByUser(usr)
}
func (dao *MySQLAccessObject) SignInUser(usr *datatypes.User) {
	dao.getValidationsByUser(usr)
	usr.IsSignedIn = true
	usr.IsValidated = true
	dao.setValidationsByUser(usr)
}
func (dao *MySQLAccessObject) SignOutUser(usr *datatypes.User) {
	dao.getValidationsByUser(usr)
	usr.IsSignedIn = false
	usr.IsValidated = false
	dao.setValidationsByUser(usr)
}
func (dao *MySQLAccessObject) UpgradeToken(tok string) bool {
	var uid string
	var email string
	logger.Debug("Validating: tok=", tok)
	dao.db.QueryRow("select uid from validations where secret = ?", tok).Scan(&uid)
	logger.Debug("Validating: uid=", uid)
	dao.db.QueryRow("select email from users where id = ?", uid).Scan(&email)
	logger.Debug("Validating: email=", email)
	usr, ok := dao.GetUserByEmail(email)
	if !ok {
		return false
	}
	dao.SignInUser(&usr)
	return true
}
