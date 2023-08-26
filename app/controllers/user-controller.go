package controllers

import (
	"encoding/json"
	"fmt"
	errorclass "interview/app/error"
	"interview/app/structs"
	"interview/app/users"
	tokenservice "interview/app/users/tokens"
	"net/http"
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

	Respond(w, r, paylaod, nil)
	return
}

func (u usercontroller) Login(w http.ResponseWriter, r *http.Request) {
	loginData := new(structs.UserLoginRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(loginData)
	if err != nil {
		Respond(w, r, nil, errorclass.NewError(errorclass.BadRequestError).Wrap(err.Error()))
	}

	err = validateLoginRequestData(*loginData)
	if err != nil {
		Respond(w, r, nil, errorclass.NewError(errorclass.BadRequestValidationError).Wrap(err.Error()))
		return
	}

	paylaod, errr := u.service.Login(r.Context(), loginData)
	if errr != nil {
		Respond(w, r, nil, errr)
		return
	}

	err = generateTokenForUserAndStoreInCache(w, paylaod.Id)
	if err != nil {
		Respond(w, r, nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error()))
	}

	Respond(w, r, paylaod, nil)
	return
}

func (u usercontroller) RefreshToken(w http.ResponseWriter, r *http.Request) {

}

func generateTokenForUserAndStoreInCache(w http.ResponseWriter, userId int) error {
	token := generateTokenForUser(userId)
	err := setTokenInRedisCache(token)
	if err != nil {
		return err
	}
	w.Header().Set("X-USER-ID", fmt.Sprint(userId))
	w.Header().Set("X-ACCESS-TOKEN", token.AccessToken)
	w.Header().Set("X-REFRESH-TOKEN", token.RefreshToken)
	return nil
}

func generateTokenForUser(userId int) tokenservice.Token {
	userClaims := tokenservice.NewUserClaims(userId)
	tokens := userClaims.GenerateToken()
	return tokens
}

func setTokenInRedisCache(token tokenservice.Token) error {
	return nil
}
