package auth

import (
	"errors"
	"os"
	"saas-estoque/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID    int64  `json:"id"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	jwt.StandardClaims
}

func GenerateToken(user *entity.User) (string, error) {

	secret := os.Getenv("JWT_SECRET_KEY")

	claims := Claims{
		UserID:    user.ID,
		Role:      user.Role,
		TokenType: "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))

}
func GenerateRefreshToken(user *entity.User) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	claims := Claims{
		UserID:    user.ID,
		Role:      user.Role,
		TokenType: "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyRefreshToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("invalid signing method")
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.TokenType != "refresh" {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("invalid signing method")
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.TokenType != "access" {
		return nil, errors.New("invalid token type")
	}
	return claims, nil
}
