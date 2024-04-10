package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (p *Auth) TokenGen(userInfo interface{}) (token string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    p.tools.InterfaceString(userInfo),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1day
	})
	token, err = claims.SignedString([]byte(p.Secret))

	return token, err

}

func (p *Auth) CheckToken(token string) (userInfo string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	tokenJwt, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.Secret), nil
	})
	if err != nil {
		return userInfo, err
	}

	claims := tokenJwt.Claims.(*jwt.StandardClaims)

	userInfo = claims.Issuer

	return
}
