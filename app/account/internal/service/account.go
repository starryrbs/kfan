package service

import (
	"context"

	"github.com/starryrbs/kfan/app/account/internal/biz"

	pb "github.com/starryrbs/kfan/api/account/service/v1"
)

func (a *AccountInterface) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	rv, err := a.ac.Repo.CreateAccount(ctx, &biz.Account{
		Name:  req.Name,
		Age:   int(req.Age),
		Email: req.Email,
		Sex:   req.Sex,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountReply{
		Id: int32(rv.Id),
	}, nil
}
func (a *AccountInterface) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountReply, error) {
	return &pb.UpdateAccountReply{}, nil
}
func (a *AccountInterface) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountReply, error) {
	return &pb.DeleteAccountReply{}, nil
}
func (a *AccountInterface) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	result, err := a.ac.Repo.GetAccount(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountReply{
		Id:    int32(result.Id),
		Name:  result.Name,
		Email: result.Email,
		Sex:   result.Sex,
		Age:   int64(result.Age),
	}, nil
}
func (a *AccountInterface) ListAccount(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountReply, error) {
	result, err := a.ac.Repo.ListAccount(ctx)
	if err != nil {
		return nil, err
	}
	rs := make([]*pb.ListAccountReply_Account, 0)
	for _, x := range result {
		rs = append(rs, &pb.ListAccountReply_Account{
			Name:  x.Name,
			Age:   int64(x.Age),
			Email: x.Email,
			Sex:   x.Sex,
		})
	}

	return &pb.ListAccountReply{
		Results: rs,
	}, nil
}
func (a *AccountInterface) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	result, err := a.ac.Repo.Login(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Id: int32(result.Id)}, nil
}
