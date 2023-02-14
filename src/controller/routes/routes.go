package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/glauberratti/MyFirstGoCRUD/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}