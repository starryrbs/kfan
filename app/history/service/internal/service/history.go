package service

import (
	"context"
	"github.com/starryrbs/kfan/app/history/service/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	pb "github.com/starryrbs/kfan/api/history/service/v1"
)

type HistoryService struct {
	pb.UnimplementedHistoryServer
	hc *biz.HistoryUseCase
}

func NewHistoryService(hc *biz.HistoryUseCase) *HistoryService {
	return &HistoryService{hc: hc}
}

func (s *HistoryService) SaveHistory(ctx context.Context, req *pb.SaveHistoryRequest) (*pb.SaveHistoryReply, error) {
	_, err := s.hc.Create(ctx, &biz.History{
		UserId:   req.UserId,
		ObjType:  req.ObjType,
		ObjId:    req.ObjId,
		CreateAt: float64(time.Now().Unix()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.SaveHistoryReply{}, nil
}
func (s *HistoryService) GetHistory(ctx context.Context, req *pb.GetHistoryRequest) (*pb.GetHistoryReply, error) {
	rvs, err := s.hc.Get(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	results := make([]*pb.GetHistoryReply_History, 0)
	for _, rv := range rvs {
		results = append(results, &pb.GetHistoryReply_History{
			ObjType:  rv.ObjType,
			ObjId:    rv.ObjId,
			CreateAt: timestamppb.New(time.Unix(int64(rv.CreateAt), 0)),
		})
	}
	return &pb.GetHistoryReply{
		Results: results,
	}, nil
}
