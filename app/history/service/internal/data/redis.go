package data

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/starryrbs/kfan/app/history/service/internal/biz"
)

func historyKey(userId int64) string {
	return fmt.Sprintf("histort-%d", userId)
}

func (h *historyRepo) GetHistoryCache(ctx context.Context, userId int64) ([]*biz.History, error) {
	results, err := h.data.rdb.ZRangeArgsWithScores(ctx, redis.ZRangeArgs{
		Key:   historyKey(userId),
		Start: 0,
		Stop:  10,
		Rev:   true,
	}).Result()
	if err != nil {
		return nil, err
	}

	histories := make([]*biz.History, 0)
	for _, z := range results {
		vs := strings.Split(z.Member.(string), "-")
		objId, err := strconv.ParseInt(vs[1], 10, 64)
		if err != nil {
			return nil, err
		}
		histories = append(histories, &biz.History{
			ObjId:    objId,
			ObjType:  vs[0],
			UserId:   userId,
			CreateAt: z.Score,
		})
	}
	return histories, nil
}

func (h historyRepo) SaveHistoryCache(ctx context.Context, history *biz.History) error {
	_, err := h.data.rdb.ZAdd(ctx, historyKey(history.UserId), &redis.Z{
		Score:  history.CreateAt,
		Member: fmt.Sprintf("%s-%d", history.ObjType, history.ObjId),
	}).Result()
	if err != nil {
		return err
	}
	return err
}
