package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type House struct {
	Id           int64
	Price        float64
	Title        string
	Community    string
	ToiletCount  int32
	KitchenCount int32
	FloorCount   int32
	HallCount    int32
	RoomCount    int32
	Image        string
}

type HouseRepo interface {
	ListHouse(ctx context.Context) ([]*House, error)
	CreateHouse(ctx context.Context, house *House) (*House, error)
	GetHouse(ctx context.Context, id int64) (*House, error)
}

type HouseUseCase struct {
	Repo HouseRepo
	log  *log.Helper
}

func NewHouseUseCase(repo HouseRepo, logger log.Logger) *HouseUseCase {
	return &HouseUseCase{
		Repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "useCase/house")),
	}
}
