package controllers

import (
	"encoding/json"
	errorclass "interview/app/error"
	"interview/app/structs"
	"interview/app/users"
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
