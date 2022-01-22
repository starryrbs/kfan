package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type History struct {
	ObjId     int64
	ObjType   string
	UserId    int64
	CreateAt  float64
	ObjDetail map[string]interface{}
	Username  string
}

type HistoryRepo interface {
	SaveHistory(context.Context, *History) (*History, error)
	GetHistory(ctx context.Context, id int64) ([]*History, error)
}

type HistoryJobRepo interface {
	PersistentSaveHistory(ctx context.Context, history *History) (*History, error)
}

type HistoryUseCase struct {
	repo HistoryRepo
	log  *log.Helper
}

func NewHistoryUseCase(repo HistoryRepo, logger log.Logger) *HistoryUseCase {
	return &HistoryUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (hc *HistoryUseCase) Create(ctx context.Context, h *History) (*History, error) {
	return hc.repo.SaveHistory(ctx, h)
}

func (hc *HistoryUseCase) Get(ctx context.Context, id int64) ([]*History, error) {
	return hc.repo.GetHistory(ctx, id)
}
