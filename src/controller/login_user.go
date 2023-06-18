package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/validation"
	"github.com/glauberratti/MyFirstGoCRUD/src/controller/model/request"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"github.com/glauberratti/MyFirstGoCRUD/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller", zap.String("jorney", "createUser"))
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("jorney", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call LoginUser service", err, zap.String("jorney", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("LoginUser controller executed successfully", zap.String("userId", domainResult.GetId()), zap.String("jorney", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
