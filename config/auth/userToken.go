package auth

import (
	"os"
	"saas-estoque/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID int64  `json:"id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(user *entity.User) (string, error) {

	secret := os.Getenv("JWT_SECRET_KEY")

	claims := Claims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))

}
