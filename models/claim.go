package models

import jwt "github.com/dgrijalva/jwt-go"

//Claim token
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
