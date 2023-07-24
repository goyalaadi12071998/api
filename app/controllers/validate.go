package controllers

import (
	"interview/app/structs"

	validation "github.com/go-ozzo/ozzo-validation"
)

func validateSignUpRequestData(data structs.UserSingupRequest) error {
	err := validation.ValidateStruct(&data,
		validation.Field(&data.Email, validation.Required),
		validation.Field(&data.Password, validation.Required, validation.Length(8, 15)),
	)

	return err
}

func validateLoginRequestData(data structs.UserLoginRequest) error {
	err := validation.ValidateStruct(&data,
		validation.Field(&data.Email, validation.Required),
		validation.Field(&data.Password, validation.Required, validation.Length(8, 15)),
	)

	return err
}
