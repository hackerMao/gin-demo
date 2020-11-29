package jwtAuth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var secretKey = []byte("never to guess!")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func GenerateToken(c *gin.Context, userId int) string {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hacker_murray",
			Subject:   "user token",
		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	fmt.Println(withClaims)
	token, err := withClaims.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return token
}

func ParseToken(token string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	withClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (i interface{}, err error) {
		return secretKey, nil
	})
	return withClaims, claims, err
}
