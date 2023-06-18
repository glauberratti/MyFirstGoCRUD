package service

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init loginUser model", zap.String("jorney", "loginUser"))
	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("LoginUser service executed successfully", zap.String("userId", user.GetId()), zap.String("jorney", "loginUser"))
	return user, token, nil
}
