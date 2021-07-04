package utils

import (
	"time"

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

func (j *JwtUtils) VerifyToken(token string) (jwt.Claims, error) {

	userClaim, err := at(time.Unix(0, 0), func() (jwt.Claims, error) {
		token, err := jwt.ParseWithClaims(token, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})
		claims, _ := token.Claims.(*UserClaim)
		// if token.Valid && ok {
		// 	return claims, err
		// }
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
	mySigningKey := []byte("AllYourBase")

	claims := UserClaim{
		*user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}
