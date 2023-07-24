package controllers

import (
	"encoding/json"
	errorclass "interview/app/error"
	"interview/app/structs"
	"interview/app/users"
	tokenservice "interview/app/users/tokens"
	"net/http"
	"time"
)

var UserController usercontroller

type usercontroller struct {
	service users.IUserService
}

func InitializeUserController(service users.IUserService) {
	UserController = usercontroller{
		service: service,
	}
}

func (u usercontroller) Signup(w http.ResponseWriter, r *http.Request) {
	signupData := new(structs.UserSingupRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(signupData)
	if err != nil {
		Respond(w, r, nil, errorclass.NewError(errorclass.BadRequestError).Wrap(err.Error()))
		return
	}

	err = validateSignUpRequestData(*signupData)
	if err != nil {
		Respond(w, r, nil, errorclass.NewError(errorclass.BadRequestValidationError).Wrap(err.Error()))
		return
	}

	paylaod, errr := u.service.Signup(r.Context(), signupData)
	if errr != nil {
		Respond(w, r, nil, errr)
		return
	}

	generateTokenForUserAndStoreInCookie(w, paylaod.Id)

	Respond(w, r, paylaod, nil)
	return
}

func generateTokenForUserAndStoreInCookie(w http.ResponseWriter, userId int) {
	tokens := generateTokenForUser(userId)
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Secure:   true,
		HttpOnly: true,
		MaxAge:   int(time.Now().Add(time.Minute * 60).Unix()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Secure:   true,
		HttpOnly: true,
		MaxAge:   int(time.Now().Add(time.Hour * 24).Unix()),
	})
}

func generateTokenForUser(userId int) tokenservice.Token {
	userClaims := tokenservice.NewUserClaims(userId)
	tokens := userClaims.GenerateToken()
	return tokens
}
