package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

/*
	Do this function to set error message in error struct.

	Args:
		err: error
		message: error message

	Returns:
		erro message

*
*/
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

/*
	Do this function to generate new hash password.

	Args:
		password: plain password

	Return:
		hashing password

*
*/
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
	Do this function to compare plain password with hash password.

	Args:
		password: plain password
		hash: hash password

	Return:
		true/false

*
*/
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
	Do this function to generate JWT token.

	Args:
		email: email user
		role: hash password

	Return:
		token JWT

*
*/
func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
