package auction

import (
	"context"
	"errors"
	"github.com/rodrigoachilles/auction-go/internal/infra/database"
	"sync"
	"testing"
	"time"

	"github.com/rodrigoachilles/auction-go/internal/entity/auction_entity"
	"github.com/rodrigoachilles/auction-go/internal/internal_error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGivenAValidParams_WhenCreateAuction_ThenShouldSaveActionSuccessfully(t *testing.T) {
	mockCollection := new(database.MockCollection)
	repo := &AuctionRepository{
		Collection:          mockCollection,
		auctionInterval:     time.Minute,
		auctionEndTimeMutex: &sync.Mutex{},
	}

	auction := &auction_entity.Auction{
		Id:          "123",
		ProductName: "Test Product",
		Category:    "Electronics",
		Description: "Test Description",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	mockCollection.On("InsertOne", mock.Anything, auctionEntityMongo, mock.Anything).
		Return(&mongo.InsertOneResult{InsertedID: "123"}, nil)

	err := repo.CreateAuction(context.Background(), auction)

	assert.Nil(t, err)
	mockCollection.AssertExpectations(t)
}

func TestGivenAValidParams_WhenCreateAuction_ThenShouldReceiveAnError(t *testing.T) {
	mockCollection := new(database.MockCollection)
	repo := &AuctionRepository{
		Collection:          mockCollection,
		auctionInterval:     time.Minute,
		auctionEndTimeMutex: &sync.Mutex{},
	}

	auction := &auction_entity.Auction{
		Id:          "123",
		ProductName: "Test Product",
		Category:    "Electronics",
		Description: "Test Description",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	mockCollection.On("InsertOne", mock.Anything, auctionEntityMongo, mock.Anything).
		Return(nil, errors.New("insert error"))

	err := repo.CreateAuction(context.Background(), auction)

	assert.NotNil(t, err)
	assert.Equal(t, internal_error.NewInternalServerError("Error trying to insert auction"), err)
	mockCollection.AssertExpectations(t)
}

func TestGivenAValidParams_WhenCreateAuction_ThenShouldCloseAutomaticallyAfterDefiniteTime(t *testing.T) {
	mockCollection := new(database.MockCollection)
	repo := &AuctionRepository{
		Collection:          mockCollection,
		auctionInterval:     time.Millisecond * 50,
		auctionEndTimeMutex: &sync.Mutex{},
	}

	auction := &auction_entity.Auction{
		Id: "123",
	}

	mockCollection.On("InsertOne", mock.Anything, mock.Anything, mock.Anything).
		Return(&mongo.InsertOneResult{InsertedID: "123"}, nil)
	mockCollection.On("UpdateByID", mock.Anything, "123", mock.MatchedBy(func(update interface{}) bool {
		updateMap, ok := update.(bson.M)
		if !ok {
			return false
		}
		setMap, ok := updateMap["$set"].(bson.M)
		if !ok {
			return false
		}
		return setMap["status"] == auction_entity.Completed
	}), mock.Anything).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil)

	err := repo.CreateAuction(context.Background(), auction)
	assert.Nil(t, err)

	// Esperar pelo tempo suficiente para que a função ChangeStatusAfter seja executada
	time.Sleep(time.Millisecond * 100)

	mockCollection.AssertExpectations(t)
}
