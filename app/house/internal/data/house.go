package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/starryrbs/kfan/app/house/internal/biz"
	"github.com/starryrbs/kfan/app/house/internal/data/ent/house"
)

type houseRepo struct {
	data *Data
	log  *log.Helper
}

func (r *houseRepo) ListHouse(ctx context.Context) ([]*biz.House, error) {
	pos, err := r.data.db.House.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.House, 0)
	for _, po := range pos {
		rv = append(rv, &biz.House{
			Id:           po.ID,
			Price:        po.Price,
			Title:        po.Title,
			Community:    po.Community,
			ToiletCount:  po.ToiletCount,
			KitchenCount: po.KitchenCount,
			FloorCount:   po.FloorCount,
			HallCount:    po.HallCount,
			RoomCount:    po.RoomCount,
			Image:        po.Image,
		})
	}
	return rv, nil
}

func (r *houseRepo) GetHouse(ctx context.Context, id int64) (*biz.House, error) {
	po, err := r.data.db.House.Query().Where(house.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.House{
		Id:           po.ID,
		Price:        po.Price,
		Title:        po.Title,
		Community:    po.Community,
		ToiletCount:  po.ToiletCount,
		KitchenCount: po.KitchenCount,
		FloorCount:   po.FloorCount,
		HallCount:    po.HallCount,
		RoomCount:    po.RoomCount,
		Image:        po.Image,
	}, nil
}

func (r *houseRepo) CreateHouse(ctx context.Context, house *biz.House) (*biz.House, error) {
	po, err := r.data.db.House.Create().
		SetCommunity(house.Community).
		SetFloorCount(house.FloorCount).
		SetHallCount(house.HallCount).
		SetKitchenCount(house.KitchenCount).
		SetPrice(0).
		SetTitle(house.Title).
		SetImage(house.Image).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.House{
		Id:           po.ID,
		Price:        po.Price,
		Title:        po.Title,
		Community:    po.Community,
		ToiletCount:  po.ToiletCount,
		KitchenCount: po.KitchenCount,
		FloorCount:   po.FloorCount,
		HallCount:    po.HallCount,
		RoomCount:    po.RoomCount,
		Image:        po.Image,
	}, nil
}

// NewHouseRepo .
func NewHouseRepo(data *Data, logger log.Logger) biz.HouseRepo {
	return &houseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
