package view

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/controller/model/response"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		Id:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
