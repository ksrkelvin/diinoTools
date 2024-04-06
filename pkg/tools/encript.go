package tools

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func (p Tools) HashPassword(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p Tools) CheckPasswordHash(password, hash string) (autorized bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (p Tools) Base64Encode(str string) (strEncoded string) {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
