package security

import (
	"github.com/arunscape/friends/commons/server/datatypes"
	"github.com/arunscape/friends/commons/server/logger"
	"github.com/arunscape/friends/commons/server/utils"

	"github.com/dgrijalva/jwt-go"

	"fmt"
	"os"
	"time"
)

var USER_SECRET = []byte(os.Getenv("TOK_SECRET"))

func CreateUserTokenLong(user datatypes.User) (string, error) {
	return createUserToken(user, time.Hour*24*7)
}
func CreateUserTokenShort(user datatypes.User) (string, error) {
	return createUserToken(user, time.Hour/2)
}
func createUserToken(user datatypes.User, duration time.Duration) (string, error) {
	logger.Debug("Created token for user: ", user.Name)
	exp := time.Now().Add(duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":        user.Name,
		"id":          user.Id,
		"email":       user.Email,
		"picture":     user.Picture,
		"groups":      user.Groups,
		"permissions": user.Permissions,
		"exp":         utils.GetMillis(exp),
	})
	tokenString, err := token.SignedString(USER_SECRET)
	return tokenString, err
}

func parseUserToken(tokStr string) (*jwt.Token, error) {
	return jwt.Parse(tokStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return USER_SECRET, nil
	})
}

func ValidateUserToken(tokStr string) bool {
	token, err := parseUserToken(tokStr)
	if err != nil {
		return false
	}
	logger.Debug("Validated token")
	mc := token.Claims.(jwt.MapClaims)
	exp := mc["exp"]
	tokenExpTime := int64(exp.(float64))
	curTime := utils.GetMillis(time.Now())
	return tokenExpTime > curTime
}

func GetUserEmailFromToken(tokStr string) string {
	token, err := parseUserToken(tokStr)
	if err != nil {
		return ""
	}
	mc := token.Claims.(jwt.MapClaims)
	return mc["email"].(string)
}

type AuthToken struct {
	Exp string `json:"exp"`

	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`

	Error string `json:"error"`
}
