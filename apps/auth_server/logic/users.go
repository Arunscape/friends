package logic

import (
	"github.com/arunscape/friends/apps/auth_server/database"
	"github.com/arunscape/friends/commons/server/datatypes"
	"github.com/arunscape/friends/commons/server/security"
	"github.com/arunscape/friends/commons/server/web_server"

	"errors"
)

// ValidateUserLogic is the logic for doing signups
func ValidateUserLogic(d interface{}, db_dat interface{}) (interface{}, error) {
	db := db_dat.(database.AccessObject)
	data := d.(*InputSignin)
	gId, name, email, picture, isValid := security.GetGoogleInfoFromToken(data.GTok)
	if !isValid {
		return nil, errors.New(web_server.USER_FAILED_TO_CREATE)
	}

	user := datatypes.User{
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

	val, err := security.MakeUserFullToken(user)
	return "{\"tok\": \"" + val + "\"}", err
}

// InputSign is the struct for both signin and signup
type InputSignin struct {
	GTok string
}
