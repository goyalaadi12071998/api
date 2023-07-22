package users

import (
	"context"
	"interview/app/models"
	"interview/app/providers/db"
)

var core usercore

type usercore struct {
	repo db.IRepo
}

func InitializeUserCore(repo db.IRepo) IUserCore {
	core = usercore{
		repo: repo,
	}
	return &core
}

func (u usercore) GetUser(ctx context.Context, filter map[string]interface{}) (*models.User, error) {
	user := &models.User{}
	err := u.repo.FindOne(user, filter)

	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (u usercore) CreateUser(ctx context.Context, userdata *models.User) (*models.User, error) {
	err := u.repo.Create(userdata)
	if err != nil {
		return nil, err
	}

	return userdata, nil
}
