package service

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("jorney", "createUser"))
	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Info("Error to call CreateUser service", zap.String("jorney", "createUser"))
		return nil, err
	}

	return userDomainRepository, nil
}
