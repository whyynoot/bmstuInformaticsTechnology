package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type Authorization struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserClaims struct {
	UserName string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}
