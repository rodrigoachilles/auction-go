package auction_usecase

import (
	"context"
	"github.com/rodrigoachilles/auction-go/internal/entity/auction_entity"
	"github.com/rodrigoachilles/auction-go/internal/entity/bid_entity"
	"github.com/rodrigoachilles/auction-go/internal/internal_error"
	"github.com/rodrigoachilles/auction-go/internal/usecase/bid_usecase"
	"time"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category"     binding:"required,min=2"`
	Description string           `json:"description"  binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition"    binding:"oneof=0 1 2"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type ProductCondition int64
type AuctionStatus int64

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

type AuctionUseCase struct {
	auctionRepository auction_entity.AuctionRepositoryInterface
	bidRepository     bid_entity.BidRepositoryInterface
}

func NewAuctionUseCase(
	auctionRepository auction_entity.AuctionRepositoryInterface,
	bidRepository bid_entity.BidRepositoryInterface) AuctionUseCaseInterface {
	return &AuctionUseCase{
		auctionRepository: auctionRepository,
		bidRepository:     bidRepository,
	}
}

type AuctionUseCaseInterface interface {
	CreateAuction(
		ctx context.Context,
		auctionInput AuctionInputDTO) (*AuctionOutputDTO, *internal_error.InternalError)

	FindAuctionById(
		ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)

	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)

	FindWinningBidByAuctionId(
		ctx context.Context,
		auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError)
}

func (au *AuctionUseCase) CreateAuction(
	ctx context.Context,
	auctionInput AuctionInputDTO) (*AuctionOutputDTO, *internal_error.InternalError) {
	auction, err := auction_entity.CreateAuction(
		auctionInput.ProductName,
		auctionInput.Category,
		auctionInput.Description,
		auction_entity.ProductCondition(auctionInput.Condition))
	if err != nil {
		return nil, err
	}

	if err := au.auctionRepository.CreateAuction(
		ctx, auction); err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}, nil
}
