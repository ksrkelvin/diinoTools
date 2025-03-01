package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT - Função para gerar o JWT
func (p *JWT) GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // O token expira após 24 horas
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "encontradev",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(p.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT - Função para validar o JWT e extrair os dados do usuário
func (p *JWT) ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse e valida o token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verifica o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return p.Secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
