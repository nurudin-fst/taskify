package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId int, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&JWTClaims{
			UserId: userId,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expiresAt),
			},
		},
	)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
