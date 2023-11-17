package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID uint) (string, error)
	VerifyToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID uint) (string, error) {
	SECRET_KEY := os.Getenv("JWT_SECRET")
	expireTime := time.Now().Add(time.Hour * 24)

	claim := jwt.MapClaims{}
	claim["userID"] = userID
	claim["exp"] = expireTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) VerifyToken(token string) (*jwt.Token, error) {
	SECRET_KEY := os.Getenv("JWT_SECRET")

	tokenVerify, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return tokenVerify, err
	}

	return tokenVerify, nil
}
