// passwords.go
package main

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is used to mask a password that is passed into the api
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash is used to check the hashed version of the password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
