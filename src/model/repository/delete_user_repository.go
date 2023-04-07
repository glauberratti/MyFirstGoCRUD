package repository

import (
	"context"
	"os"

	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/logger"
	"github.com/glauberratti/MyFirstGoCRUD/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser repository", zap.String("jorney", "deleteUser"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user", err, zap.String("jorney", "deleteUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("DeleteUser repository executed successfully", zap.String("userId", userId), zap.String("jorney", "deleteUser"))
	return nil
}
