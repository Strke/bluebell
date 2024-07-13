package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("开发")
var KeyFunc jwt.Keyfunc = func(token *jwt.Token) (i interface{}, err error) { return MySecret, nil }

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userID int64) (aToken, rToken string, err error) {
	c := MyClaims{
		userID,
		"username",
		jwt.StandardClaims{
			//(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour)
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	aToken, err = Token.SignedString(MySecret)

	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "bluebell",
	}).SignedString(MySecret)
	return
}

func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)

	token, err = jwt.ParseWithClaims(tokenString, claims, KeyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		err = errors.New("invalid token")
	}
	return
}

func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	if _, err = jwt.Parse(rToken, KeyFunc); err != nil {
		return
	}

	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, KeyFunc)
	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	return
}
