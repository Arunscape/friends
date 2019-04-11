package logic

import (
	"github.com/arunscape/friends/apps/api_server/database"

	"github.com/dgrijalva/jwt-go"

	"fmt"
	"os"
)

var USER_SECRET = []byte(os.Getenv("TOK_SECRET"))

func MakeUserFullToken(user database.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"id":   user.Id,
	})
	tokenString, err := token.SignedString(USER_SECRET)
	return tokenString, err
}

func ValidateUserToken(tokStr string) bool {
	token, err := jwt.Parse(tokStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return USER_SECRET, nil
	})
	if err != nil {
		return false
	}
	mc := token.Claims.(jwt.MapClaims)
	exp := mc["exp"]
	if exp == nil {
		return false
	}
	return true

}

func GetGoogleIdFromToken(gIdTok string) (string, bool) {
	token, err := jwt.Parse(gIdTok, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return USER_SECRET, nil
	})
	if err != nil {
		return "", false
	}
	mc := token.Claims.(jwt.MapClaims)
	sub := mc["sub"]
	if sub == nil {
		return "", false
	}
	return sub.(string), true
}