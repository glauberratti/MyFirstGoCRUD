package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"github.com/glauberratti/MyFirstGoCRUD/src/model"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/repository/entity"
	"github.com/glauberratti/MyFirstGoCRUD/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository", zap.String("jorney", "findUserByEmail"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("jorney", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error trying to find user by email")
		logger.Error(errorMessage, err, zap.String("jorney", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserByEmail repository executed successfully",
		zap.String("jorney", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.Id.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById repository", zap.String("jorney", "findUserById"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("jorney", "findUserById"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := fmt.Sprintf("Error trying to find user by id")
		logger.Error(errorMessage, err, zap.String("jorney", "findUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info(
		"FindUserById repository executed successfully",
		zap.String("jorney", "findUserById"),
		zap.String("userId", userEntity.Id.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}
