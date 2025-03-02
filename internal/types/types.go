package types

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Uid string  `json:"uid"`
	Exp float64 `json:"exp"`
	jwt.RegisteredClaims
}

type UserCtxKey string

var AuthKey UserCtxKey = "foodgo-auth"
