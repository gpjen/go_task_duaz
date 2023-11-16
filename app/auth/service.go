package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
	VerifyToken(token string) (*jwt.Token, error)
}

// type
