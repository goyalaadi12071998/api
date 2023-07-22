package controllers

import (
	"context"
	"encoding/json"
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

	cookie, err := generateSessionForUser(r.Context(), paylaod.Id)
	if err != nil {
		Respond(w, r, paylaod, nil)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    cookie,
		Secure:   true,
		HttpOnly: true,
	})

	Respond(w, r, paylaod, nil)
	return
}

func generateSessionForUser(ctx context.Context, userId int) (string, error) {
	userClaims := tokenservice.NewUserClaims(userId)
	cookieId, err := userClaims.GenerateTokenAndStoreInCache(ctx)
	if err != nil {
		return "", err
	}

	return cookieId, nil
}
