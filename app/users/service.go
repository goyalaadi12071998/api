package users

var service userservice

type userservice struct {
	core IUserCore
}

func InitializeUserService(core IUserCore) IUserService {
	service = userservice{
		core: core,
	}
	return &service
}
