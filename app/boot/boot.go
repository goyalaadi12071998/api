package boot

import (
	"context"
	"interview/app/common"
	"interview/app/controllers"
	"interview/app/providers/db"
	redisclient "interview/app/providers/redis"
	"interview/app/router"
	"interview/app/users"
)

func initMySqlClient() error {
	config := common.GetConfig().Database
	err := db.InitDB(db.Config{
		Name:     config.Name,
		Host:     config.Host,
		Port:     config.Port,
		Username: config.Username,
		Password: config.Password,
	})

	return err
}

func initRedisClient(ctx context.Context) error {
	config := common.GetConfig().Redis
	err := redisclient.InitRedisClient(ctx, redisclient.Config{
		Host: config.Host,
		Port: config.Port,
	})

	return err
}

func initProviders(ctx context.Context) error {
	err := initMySqlClient()
	if err != nil {
		return err
	}

	err = initRedisClient(ctx)
	if err != nil {
		return err
	}

	return nil
}

func initRoutes() error {
	configs := common.GetConfig()
	err := router.InitializeRouter(router.CoreConfigs{
		Name: configs.Core.Name,
		Host: configs.Core.Host,
		Port: configs.Core.Port,
	})
	return err
}

func initControllers() error {
	database := db.GetDb()
	repo := db.InitializeRepo(database)

	usercore := users.InitializeUserCore(repo)
	userservice := users.InitializeUserService(usercore)

	controllers.InitializeAppController()
	controllers.InitializeUserController(userservice)

	return nil
}

func Init(ctx context.Context, env string) error {
	err := common.InitConfig(env)
	if err != nil {
		return err
	}

	err = initProviders(ctx)
	if err != nil {
		return err
	}

	err = initControllers()
	if err != nil {
		return err
	}

	err = initRoutes()
	if err != nil {
		return err
	}

	return nil
}
