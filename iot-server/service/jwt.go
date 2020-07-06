package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("hui_iot(*^*)")

type Claims struct {
	AppKey string `json:"appKey"`
	jwt.StandardClaims
}

func GenerateToken(appKey string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)

	claims := Claims{
		appKey,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hui-iot",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
