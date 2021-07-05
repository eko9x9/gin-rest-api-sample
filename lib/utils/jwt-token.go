package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/eko9x9/gin-rest-api-sample/config"
	"github.com/golang-jwt/jwt"
)

type JwtUtils struct {
}

type UserClaim struct {
	UserInfo
	jwt.StandardClaims
}

type UserInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func (j *JwtUtils) ValidateToken(token string) (jwt.Claims, error) {
	key, err := config.GetJwtKey()
	if err != nil {
		log.Fatalf("Error get key of jwt: %v", err)
	}

	userClaim, err := at(time.Unix(0, 0), func() (jwt.Claims, error) {
		token, err := jwt.ParseWithClaims(token, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
		claims, ok := token.Claims.(*UserClaim)

		if !token.Valid || !ok {
			fmt.Println(token.Valid, ok)
		}
		return claims, err
	})

	return userClaim, err
}

func at(t time.Time, f func() (jwt.Claims, error)) (jwt.Claims, error) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	claim, err := f()
	jwt.TimeFunc = time.Now

	return claim, err
}

func (j *JwtUtils) GenerateToken(user *UserInfo) (string, error) {
	key, err := config.GetJwtKey()
	if err != nil {
		log.Fatalf("Error get key of jwt: %v", err)
	}

	claims := UserClaim{
		*user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)

	return ss, err
}

type JwtUserRoles int

const (
	Admin JwtUserRoles = iota + 1
	User
	None
)
