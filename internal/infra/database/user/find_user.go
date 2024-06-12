package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/rodrigoachilles/auction-go/configuration/logger"
	"github.com/rodrigoachilles/auction-go/internal/entity/user_entity"
	"github.com/rodrigoachilles/auction-go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *UserRepository) FindUsers(
	ctx context.Context) ([]user_entity.User, *internal_error.InternalError) {
	cursor, err := ur.Collection.Find(ctx, bson.M{})
	if err != nil {
		logger.Error("Error finding users", err)
		return nil, internal_error.NewInternalServerError("Error finding users")
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	var usersMongo []UserEntityMongo
	if err := cursor.All(ctx, &usersMongo); err != nil {
		logger.Error("Error decoding users", err)
		return nil, internal_error.NewInternalServerError("Error decoding users")
	}

	var usersEntity []user_entity.User
	for _, user := range usersMongo {
		usersEntity = append(usersEntity, user_entity.User{
			Id:   user.Id,
			Name: user.Name,
		})
	}

	return usersEntity, nil
}

func (ur *UserRepository) FindUserById(
	ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": userId}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with this id = %s", userId), err)
			return nil, internal_error.NewNotFoundError(
				fmt.Sprintf("User not found with this id = %s", userId))
		}

		logger.Error("Error trying to find user by userId", err)
		return nil, internal_error.NewInternalServerError("Error trying to find user by userId")
	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil
}
