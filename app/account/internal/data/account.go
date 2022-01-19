package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/starryrbs/kfan/app/account/internal/biz"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func (a *accountRepo) CreateAccount(ctx context.Context, account *biz.Account) (*biz.Account, error) {
	po, err := a.data.db.Account.
		Create().
		SetName(account.Name).
		SetAge(account.Age).
		SetEmail(account.Email).
		SetSex(account.Sex).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Account{
		Id:    po.ID,
		Name:  po.Name,
		Email: po.Email,
		Age:   po.Age,
		Sex:   po.Sex,
	}, nil
}

func (a *accountRepo) ListAccount(ctx context.Context) ([]*biz.Account, error) {
	pos, err := a.data.db.Account.Query().All(ctx)
	if err != nil {
		return nil, nil
	}
	rv := make([]*biz.Account, 0)
	for _, po := range pos {
		rv = append(rv, &biz.Account{
			Id:    po.ID,
			Name:  po.Name,
			Email: po.Email,
			Age:   po.Age,
			Sex:   po.Sex,
		})
	}
	return rv, nil
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}