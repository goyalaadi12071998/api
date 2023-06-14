package users

import "interview/app/providers/db"

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
