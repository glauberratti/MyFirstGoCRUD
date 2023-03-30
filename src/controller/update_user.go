package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/validation"
	"github.com/glauberratti/MyFirstGoCRUD/src/controller/model/request"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller", zap.String("jorney", "updateUser"))
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("jorney", "updateUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("User id is not a valid hex", err, zap.String("jorney", "updateUser"))
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser service", err, zap.String("jorney", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully", zap.String("userId", userId), zap.String("jorney", "updateUser"))

	c.Status(http.StatusOK)
}
