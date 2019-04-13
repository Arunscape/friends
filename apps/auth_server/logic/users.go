package logic

import (
	"github.com/arunscape/friends/commons/server/web_server"
	"github.com/arunscape/friends/apps/auth_server/database"

	"errors"
)


// ValidateUserLogic is the logic for doing signups
func ValidateUserLogic(d interface{}, db_dat interface{}) (interface{}, error) {
    db := db_dat.(database.AccessObject)
	data := d.(*InputSignin)
	gId, name, email, picture, isValid := GetGoogleInfoFromToken(data.GTok)
	if !isValid {
		return nil, errors.New(web_server.USER_FAILED_TO_CREATE)
	}

	user := database.User{
		AuthId:  gId,
		Name:    name,
		Email:   email,
		Picture: picture,
	}
	db.CreateNewUser(user)
	user, found := db.GetUserByAuthId(gId)
	if !found {
		return nil, errors.New(web_server.USER_FAILED_TO_CREATE)
	}

	val, err := MakeUserFullToken(user)
	return "{\"tok\": \"" + val + "\"}", err
}

// InputSign is the struct for both signin and signup
type InputSignin struct {
	GTok string
}
