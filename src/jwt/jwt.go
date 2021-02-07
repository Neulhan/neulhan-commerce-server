package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"neulhan-commerce-server/src/config"
	"time"
)

type Claims struct {
	UserID int
	jwt.StandardClaims
}

var expirationTime = time.Minute * 60

func CreateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(expirationTime)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetEnv("SECRET_KEY")))

	if err != nil {
		return "", fmt.Errorf("token signed error")
	}
	return tokenString, nil
}

func ParseToken(myToken string) (*Claims, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(myToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("SECRET_KEY")), nil
	})

	if err != nil {
		return claims, err
	}

	return claims, nil
}
