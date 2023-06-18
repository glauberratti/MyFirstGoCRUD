package service

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById services", zap.String("jorney", "findUserById"))

	return ud.userRepository.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("jorney", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("jorney", "findUserByEmail"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
