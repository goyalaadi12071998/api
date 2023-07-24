package tokenservice

import (
	"interview/app/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

type userClaims struct {
	id int
	jwt.StandardClaims
}

type Token struct {
	AccessToken  string
	RefreshToken string
}

type ITokensService interface {
	newAccessToken(duration time.Duration) (string, error)
	newRefreshToken(duration time.Duration) (string, error)
	GenerateToken() Token
}

func (u userClaims) newAccessToken(duration time.Duration) (string, error) {
	u.StandardClaims = jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(duration).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return accessToken.SignedString([]byte("thisisasecret"))
}

func (u userClaims) newRefreshToken(duration time.Duration) (string, error) {
	u.StandardClaims = jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(duration).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return refreshToken.SignedString([]byte("thisisasecret"))
}

func (u userClaims) GenerateToken() Token {
	accessToken, _ := u.newAccessToken(constants.JWT_ACCESS_TOKEN_TIME_DURATION)
	refreshToken, _ := u.newRefreshToken(constants.JWT_REFRESH_TOKEN_TIME_DURATION)

	return Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func NewUserClaims(id int) ITokensService {
	return userClaims{
		id: id,
	}
}
