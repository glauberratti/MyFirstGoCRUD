package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/database/mongodb"
	"github.com/glauberratti/MyFirstGoCRUD/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":5000"); err != nil {
		log.Fatal(err)
		return
	}
}
