package service

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("jorney", "createUser"))
	userDomain.EncryptPassword()

	user, err := ud.FindUserByEmailServices(userDomain.GetEmail())
	if err != nil {
		logger.Info("Error to call FindUserByEmailServices service", zap.String("jorney", "createUser"))
		return nil, err
	}

	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Info("Error to call CreateUser service", zap.String("jorney", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully", zap.String("userId", userDomainRepository.GetId()), zap.String("jorney", "createUser"))
	return userDomainRepository, nil
}
