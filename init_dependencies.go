package main

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/controller"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/repository"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepository)
	return controller.NewUserControllerInterface(userService)
}
