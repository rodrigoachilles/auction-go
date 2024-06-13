package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodrigoachilles/auction-go/configuration/database/mongodb"
	"github.com/rodrigoachilles/auction-go/configuration/logger"
	"github.com/rodrigoachilles/auction-go/internal/infra/api/web/controller/auction_controller"
	"github.com/rodrigoachilles/auction-go/internal/infra/api/web/controller/bid_controller"
	"github.com/rodrigoachilles/auction-go/internal/infra/api/web/controller/user_controller"
	"github.com/rodrigoachilles/auction-go/internal/infra/database/auction"
	"github.com/rodrigoachilles/auction-go/internal/infra/database/bid"
	"github.com/rodrigoachilles/auction-go/internal/infra/database/user"
	"github.com/rodrigoachilles/auction-go/internal/usecase/auction_usecase"
	"github.com/rodrigoachilles/auction-go/internal/usecase/bid_usecase"
	"github.com/rodrigoachilles/auction-go/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionController := initDependencies(databaseConnection)

	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.FindUsers)
	router.GET("/users/:id", userController.FindUserById)

	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auctions", auctionController.FindAuctions)
	router.GET("/auctions/:id", auctionController.FindAuctionById)
	router.GET("/auctions/:id/winner", auctionController.FindWinningBidByAuctionId)
	router.GET("/auctions/:id/bids", bidController.FindBidByAuctionId)

	router.POST("/bids", bidController.CreateBid)

	serverPort := ":8080"
	logger.Info(fmt.Sprintf("Starting server on port %s ...", serverPort[1:]))

	_ = router.Run(serverPort)
}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController) {

	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(
		user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(
		auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return
}
