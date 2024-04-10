package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var signingKey = []byte(viper.GetString("jwt.signingKey"))

type JwtCustClaims struct {
	Id        int64
	NickyName string
	jwt.RegisteredClaims
}

func GenerateToken(id int64, nickyName string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		Id:        id,
		NickyName: nickyName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString(signingKey)
}
func ParseToken(tokenStr string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err == nil && !token.Valid {
		err = errors.New("invalid Token")
	}
	return iJwtCustClaims, err
}
func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}
