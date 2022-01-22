package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/starryrbs/kfan/app/house/internal/biz"

	pb "github.com/starryrbs/kfan/api/house/service/v1"
)

type HouseService struct {
	pb.UnimplementedHouseServer
	hc  *biz.HouseUseCase
	log *log.Helper
}

func NewHouseService(hc *biz.HouseUseCase, logger log.Logger) *HouseService {
	return &HouseService{
		hc:  hc,
		log: log.NewHelper(log.With(logger, "service/interface")),
	}
}

func (s *HouseService) CreateHouse(ctx context.Context, req *pb.CreateHouseRequest) (*pb.CreateHouseReply, error) {
	house := &biz.House{
		Title:        req.Title,
		Community:    req.Community,
		ToiletCount:  req.ToiletCount,
		KitchenCount: req.KitchenCount,
		FloorCount:   req.FloorCount,
		HallCount:    req.HallCount,
		RoomCount:    req.RoomCount,
		Image:        req.Image,
	}

	rv, err := s.hc.Repo.CreateHouse(ctx, house)
	if err != nil {
		return nil, err
	}
	return &pb.CreateHouseReply{
		Id: rv.Id,
	}, nil
}
func (s *HouseService) GetHouse(ctx context.Context, req *pb.GetHouseRequest) (*pb.GetHouseReply, error) {
	rv, err := s.hc.Repo.GetHouse(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetHouseReply{
		Id:           rv.Id,
		Title:        rv.Title,
		Community:    rv.Community,
		ToiletCount:  rv.ToiletCount,
		KitchenCount: rv.KitchenCount,
		FloorCount:   rv.FloorCount,
		HallCount:    rv.HallCount,
		RoomCount:    rv.RoomCount,
	}, nil
}
func (s *HouseService) ListHouse(ctx context.Context, req *pb.ListHouseRequest) (*pb.ListHouseReply, error) {

	rvs, err := s.hc.Repo.ListHouse(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*pb.ListHouseReply_House, 0)

	for _, rv := range rvs {
		results = append(results, &pb.ListHouseReply_House{
			Id:        rv.Id,
			Title:     rv.Title,
			Community: rv.Community,
			Description: fmt.Sprintf("%d室%d厅%d卫%d厨",
				rv.RoomCount, rv.HallCount, rv.ToiletCount, rv.KitchenCount),
			Image: rv.Image,
		})
	}

	return &pb.ListHouseReply{
		Results: results,
	}, nil
}
