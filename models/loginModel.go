package models

import "github.com/golang-jwt/jwt"

//model login user
type Login struct {
	UserName string
	Password string
}

//Claim
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}