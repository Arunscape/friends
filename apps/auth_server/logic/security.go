package logic

import (
	"github.com/arunscape/friends/apps/auth_server/database"

	"github.com/dgrijalva/jwt-go"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var USER_SECRET = []byte(os.Getenv("TOK_SECRET"))

func MakeUserFullToken(user database.User) (string, error) {
	exp := time.Now().Add(time.Minute * 5)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"id":   user.Id,
		"exp":  exp.Unix()*1000 + int64(exp.Nanosecond()/1000000),
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
	expTime := int64(exp.(float64))
	t := time.Now()
	curTime := t.Unix()*1000 + int64(t.Nanosecond()/1000000)
	return expTime > curTime

}

type GoogleAuth struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	Aud string `json:"aud"`
	Iat string `json:"iat"`
	Exp string `json:"exp"`

	Email          string `json:"email"`
	Email_verified string `json:"email_verified"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Locale         string `json:"locale"`

	Error string `json:"error"`
}

func validateGoogleToken(gIdTok string) (GoogleAuth, bool) {
	// Make http request to https://oauth2.googleapis.com/tokeninfo?id_token=XYZ123
	resp, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + gIdTok)
	if err != nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return GoogleAuth{}, false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var gAuth GoogleAuth
	json.Unmarshal(body, &gAuth)
	if gAuth.Error == "" {
		return gAuth, true
	}
	return gAuth, false
}

func GetGoogleIdFromToken(gIdTok string) (string, bool) {
	tok, isGood := validateGoogleToken(gIdTok)
	return tok.Sub, isGood
}

func GetGoogleInfoFromToken(gIdTok string) (string, string, string, string, bool) {
	tok, isGood := validateGoogleToken(gIdTok)
	return tok.Sub, tok.Name, tok.Email, tok.Picture, isGood
}
