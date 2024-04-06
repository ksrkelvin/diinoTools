package tools

import (
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
