package tokenservice

import (
	"context"
	"encoding/json"
	redisclient "interview/app/providers/redis"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type userClaims struct {
	id int
	jwt.StandardClaims
}

type Token struct {
	accessToken  string
	refreshToken string
}

type ITokensService interface {
	newAccessToken(duration time.Duration) (string, error)
	newRefreshToken(duration time.Duration) (string, error)
	generateToken() Token
	GenerateTokenAndStoreInCache(ctx context.Context) (string, error)
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

func (u userClaims) generateToken() Token {
	accessToken, _ := u.newAccessToken(time.Duration(time.Minute * 60))
	refreshToken, _ := u.newRefreshToken(time.Duration(time.Hour * 24))

	return Token{
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

func (u userClaims) GenerateTokenAndStoreInCache(ctx context.Context) (string, error) {
	client := redisclient.GetClient()
	token := u.generateToken()
	cookie, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	cookieId := cookie.String()
	tokenData, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	err = client.Set(ctx, cookieId, tokenData, time.Duration(time.Minute*60))
	if err != nil {
		return "", err
	}

	return cookieId, nil
}

func NewUserClaims(id int) ITokensService {
	return userClaims{
		id: id,
	}
}
