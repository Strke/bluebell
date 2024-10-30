package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Token过期时间

const TokenExpireDuration = time.Hour * 2

//设置秘钥

var MySecret = []byte("开发")
var KeyFunc jwt.Keyfunc = func(token *jwt.Token) (i interface{}, err error) { return MySecret, nil }

// 自定义声明结构体并且内嵌jwt.StandardClaims
// jwt.StandardClaims内部只包含官方字段
// 由于我们要额外添加字段，所以创建了一个新的结构体
// 在新的结构体里面我们加上了用户ID和用户名

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userID int64) (aToken, rToken string, err error) {
	c := MyClaims{
		userID,     //自定义字段，用户ID
		"username", //自定义字段，用户名
		jwt.StandardClaims{

			//(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour)

			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //官方设定字段，过期时间
			Issuer:    "bluebell",                                 //官方设定字段，签发人
		},
	}
	// 使用指定加密方式HS256生成签名Token
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用MySecret加密并获得accessToken
	aToken, err = Token.SignedString(MySecret)
	// 获得一个refreshToken
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		Issuer:    "bluebell",
	}).SignedString(MySecret)
	return
}

func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	// 解析Token
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
