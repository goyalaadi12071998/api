package users

import (
	"context"
	errorclass "interview/app/error"
	"interview/app/models"
	"interview/app/structs"
)

type IUserService interface {
	Signup(ctx context.Context, data *structs.UserSingupRequest) (*structs.UserSingupRequestResponse, *errorclass.Error)
	Login(ctx context.Context, data *structs.UserLoginRequest) (*structs.UserLoginRequestResponse, *errorclass.Error)
}

type IUserCore interface {
	GetUser(ctx context.Context, filter map[string]interface{}) (*models.User, error)

	CreateUser(ctx context.Context, userdata *models.User) (*models.User, error)
}
