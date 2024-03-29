package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller", zap.String("jorney", "deleteUser"))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("User id is not a valid hex", err, zap.String("jorney", "deleteUser"))
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call DeleteUser service", err, zap.String("jorney", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully", zap.String("userId", userId), zap.String("jorney", "deleteUser"))

	c.Status(http.StatusOK)
}
