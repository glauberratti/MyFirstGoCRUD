package service

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model", zap.String("jorney", "updateUser"))

	err := ud.userRepository.UpdateUser(id, userDomain)
	if err != nil {
		logger.Info("Error to call UpdateUser service", zap.String("jorney", "updateUser"))
		return err
	}

	logger.Info("UpdateUser service executed successfully", zap.String("userId", id), zap.String("jorney", "updateUser"))
	return nil
}
