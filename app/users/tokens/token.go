package tokenservice

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type userClaims struct {
	id int
	jwt.StandardClaims
}

type ITokensService interface {
	NewAccessToken() (string, error)
	NewRefreshToken() (string, error)
}

func (u userClaims) NewAccessToken(duration time.Duration) (string, error) {
	u.StandardClaims = jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(duration).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return accessToken.SignedString([]byte("thisisasecret"))
}

func (u userClaims) NewRefreshToken(duration time.Duration) (string, error) {
	u.StandardClaims = jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(duration).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return refreshToken.SignedString([]byte("thisisasecret"))
}

func NewUserClaims(id int) userClaims {
	return userClaims{
		id: id,
	}
}
