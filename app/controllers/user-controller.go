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
	data := new(structs.UserSingupRequest)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		Respond(w, r, nil, errorclass.NewError(errorclass.BadRequestError).Wrap("input data is not valid"))
	}

	Respond(w, r, data, nil)
}
