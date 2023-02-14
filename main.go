package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/database/mongodb"
	"github.com/glauberratti/MyFirstGoCRUD/src/controller"
	"github.com/glauberratti/MyFirstGoCRUD/src/controller/routes"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/repository"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		panic(err)
	}
	println(db)

	router := gin.Default()
	//Init dependecies

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepository)
	userController := controller.NewUserControllerInterface(userService)

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}
