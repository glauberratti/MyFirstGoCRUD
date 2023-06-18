package service

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository: userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIdServices(id string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
