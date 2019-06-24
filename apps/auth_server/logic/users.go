package logic

import (
	"github.com/arunscape/friends/apps/auth_server/database"
	"github.com/arunscape/friends/commons/server/datatypes"
	"github.com/arunscape/friends/commons/server/logger"
	"github.com/arunscape/friends/commons/server/mail"
	"github.com/arunscape/friends/commons/server/security"
	"github.com/arunscape/friends/commons/server/utils"
	"github.com/arunscape/friends/commons/server/web_server"

	"errors"
	"os"
)

func UpgradeLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*Token)
	logger.Debug(data)

	usr, err := getUserFromToken(db, data.Tok)
	if err != nil {
		return usr, err
	}
	if !usr.IsSignedIn {
		return nil, errors.New(web_server.TOKEN_FORBIDDEN)
	}
	db.SignInUser(&usr)
	tok, err := security.CreateUserTokenLong(usr)

	logger.Info("Validating: ", usr.Email)
	return Token{tok}, err
}

// SettingsLogic Sets the users settings based on the token
func SettingsLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*Settings)
	usr, err := getUserFromToken(db, data.Tok)
	if err != nil {
		return nil, errors.New(web_server.USER_NOT_FOUND)
	}
	usr.Settings = data.Settings
	db.SaveUserSettings(usr)
	tok, err := security.CreateUserTokenLong(usr)
	return &Token{tok}, nil
}

// SigninLogic creates a short lived token, sends email link, then returns token, sets isValidated to false and isSignedIn to true
func SigninLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*Email)
	return startSigninProcess(data.Email, db)
}

// SignupLogic adds name, email, pic, etc to DB, then acts as if signin
func SignupLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*Signup)
	logger.Debug("Creating new User: ", data)
	db.CreateNewUser(&datatypes.User{Name: data.Name, Email: data.Email, Picture: data.Pic})
	return startSigninProcess(data.Email, db)
}

// IsUserLogic check if the user exists
func IsUserLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*Email)
	_, isUser := db.GetUserByEmail(data.Email)
	logger.Debug("Did I find User? ", isUser)
	if !isUser {
		return nil, errors.New(web_server.USER_NOT_FOUND)
	}
	return nil, nil
}

func SignoutLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*Token)

	usr, err := getUserFromToken(db, data.Tok)
	if err != nil {
		return usr, err
	}

	db.SignOutUser(&usr)
	return nil, nil
}

func startSigninProcess(email string, db database.AccessObject) (interface{}, error) {
	secret := utils.UUID()
	usr, isUser := db.GetUserByEmail(email)
	if !isUser {
		return nil, errors.New(web_server.USER_DOES_NOT_EXIST)
	}
	db.AddUserValidation(&usr, secret)
	link := "http://auth." + os.Getenv("DOMAIN") + "/validate/" + secret // TODO: email link
	mail.SendEmail([]string{usr.Email}, "Your signup link for friends", link)
	logger.Info("Created secure link: ", link)
	tok, err := security.CreateUserTokenShort(usr)
	if err != nil {
		return nil, errors.New(web_server.UNKNOWN)
	}
	return Token{Tok: tok}, nil
}

func getUserFromToken(db database.AccessObject, tokStr string) (datatypes.User, error) {
	if !security.ValidateUserToken(tokStr) {
		return datatypes.User{}, errors.New(web_server.TOKEN_FORBIDDEN)
	}
	email := security.GetUserEmailFromToken(tokStr)
	usr, isReal := db.GetUserByEmail(email)
	if !isReal {
		return datatypes.User{}, errors.New(web_server.USER_NOT_FOUND)
	}
	return usr, nil
}

type (
	Token struct {
		Tok string
	}
	Settings struct {
		Tok      string
		Settings string
	}
	Email struct {
		Email string
	}
	Signup struct {
		Email string
		Name  string
		Pic   string
	}
)
