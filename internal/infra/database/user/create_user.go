package user

import (
	"context"
	"github.com/rodrigoachilles/auction-go/configuration/logger"
	"github.com/rodrigoachilles/auction-go/internal/entity/user_entity"
	"github.com/rodrigoachilles/auction-go/internal/internal_error"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) CreateUser(
	ctx context.Context,
	userEntity *user_entity.User) *internal_error.InternalError {
	userEntityMongo := &UserEntityMongo{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}
	_, err := ur.Collection.InsertOne(ctx, userEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert user", err)
		return internal_error.NewInternalServerError("Error trying to insert user")
	}

	return nil
}
