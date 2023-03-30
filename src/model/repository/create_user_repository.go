package repository

import (
	"context"
	"os"

	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("jorney", "createUser"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user", err, zap.String("jorney", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)
	logger.Info("CreateUser repository executed successfully", zap.String("userId", value.Id.Hex()), zap.String("jorney", "createUser"))
	return converter.ConvertEntityToDomain(*value), nil
}
