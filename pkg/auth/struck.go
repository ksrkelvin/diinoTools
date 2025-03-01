package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// JWT - Estrutura  jwt
type JWT struct {
	Secret []byte
}

// Claims - Struct para representar os dados do usuário no token
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Init - Inicializa a autenticação
func Init(secret string) (auth *JWT, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	newAuth := &JWT{
		Secret: []byte(secret),
	}
	return newAuth, err
}
