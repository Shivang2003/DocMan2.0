package types

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
