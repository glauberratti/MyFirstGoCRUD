package converter

import (
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Name:     domain.GetName(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}
