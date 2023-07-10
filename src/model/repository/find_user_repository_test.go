package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/glauberratti/MyFirstGoCRUD/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	err := os.Setenv("MONGODB_USER_COLLECTION", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			Id:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetId(), userEntity.Id.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomain.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, userDomain.GetPassword(), userDomain.GetPassword())
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserById(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	err := os.Setenv("MONGODB_USER_COLLECTION", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_id_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			Id:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById(userEntity.Id.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetId(), userEntity.Id.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomain.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, userDomain.GetPassword(), userDomain.GetPassword())
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("test")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("test")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, fmt.Sprintf("User not found with this id: test"))
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	err := os.Setenv("MONGODB_USER_COLLECTION", collection_name)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_email_and_password_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			Id:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetId(), userEntity.Id.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomain.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, userDomain.GetPassword(), userDomain.GetPassword())
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("test", "testpass")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmailAndPassword("test", "testpass")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.Id},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
