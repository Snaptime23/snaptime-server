package jwt

import (
	"errors"
	"github.com/Snaptime23/snaptime-server/v2/tools/errno"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecret = []byte("snaptime_server_998244353_1000000007_114514_1919810_snaptime_server")

type MyClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func GenToken(userid string, username string) (aToken string, err error) {
	aclaims := MyClaims{
		userid,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
			Issuer:    "Park-Jiyeon",
		},
	}

	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, aclaims).SignedString(mySecret)
	return
}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			err = errno.NewErrNo("Token解析出错" + jwt.ErrTokenExpired.Error())
		}
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
