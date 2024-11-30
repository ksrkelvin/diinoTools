package tools

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword - Hash a password
func (p Tools) HashPassword(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - Check if the password is correct
func (p Tools) CheckPasswordHash(password, hash string) (autorized bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Base64Encode - Encode a string to base64
func (p Tools) Base64Encode(str string) (strEncoded string) {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// HashString - Hash a string
func (p Tools) HashString(str string) (strHash string) {
	strHash = hash([]byte(str))
	return strHash
}

// hash - Hash a byte array into a string
func hash(b []byte) (hash string) {
	h := sha1.New()
	h.Write(b)
	sum := h.Sum(nil)
	hash = fmt.Sprintf("%x", sum)
	return hash
}
