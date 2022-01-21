package data

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/starryrbs/kfan/app/history/service/internal/biz"
	"time"
)

type historyRepo struct {
	data *Data
	log  *log.Helper
}

type historyJobRepo struct {
	data *JobData
	log  *log.Helper
}

func (h historyJobRepo) PersistentSaveHistory(ctx context.Context, history *biz.History) (*biz.History, error) {
	_, err := h.data.db.History.Create().
		SetObjID(history.ObjId).
		SetObjType(history.ObjType).
		SetUserID(history.UserId).
		SetCreatedAt(time.Unix(int64(history.CreateAt), 0)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return history, nil
}

type HistoryMessage struct{}

func (h historyRepo) SaveHistory(ctx context.Context, history *biz.History) (*biz.History, error) {
	err := h.SaveHistoryCache(ctx, history)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(history)
	// 向kafka写入
	h.data.kp.Input() <- &sarama.ProducerMessage{
		Topic: "history",
		Value: sarama.ByteEncoder(b),
	}
	return history, nil
}

func (h historyRepo) GetHistory(ctx context.Context, id int64) ([]*biz.History, error) {
	return h.GetHistoryCache(ctx, id)
}

// NewHistoryRepo .
func NewHistoryRepo(data *Data, logger log.Logger) biz.HistoryRepo {
	return &historyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewHistoryJobRepo(data *JobData, logger log.Logger) biz.HistoryJobRepo {
	return &historyJobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
