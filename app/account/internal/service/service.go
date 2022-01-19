package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/starryrbs/kfan/api/account/service/v1"
	"github.com/starryrbs/kfan/app/account/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAccountInterface)

type AccountInterface struct {
	pb.UnimplementedAccountServer
	ac  *biz.AccountUseCase
	log *log.Helper
}

func NewAccountInterface(uc *biz.AccountUseCase, logger log.Logger) *AccountInterface {
	return &AccountInterface{ac: uc, log: log.NewHelper(log.With(logger, "service/interface"))}
}
