package service

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser model", zap.String("jorney", "deleteUser"))

	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Info("Error to call DeleteUser service", zap.String("jorney", "deleteUser"))
		return err
	}

	logger.Info("DeleteUser service executed successfully", zap.String("userId", id), zap.String("jorney", "deleteUser"))
	return nil
}
