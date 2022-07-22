package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/starryrbs/kfan/app/account/internal/biz"
	"github.com/starryrbs/kfan/app/account/internal/data/ent/account"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func (a *accountRepo) GetAccount(ctx context.Context, id int) (*biz.Account, error) {
	po, err := a.data.db.Account.Query().Where(account.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	var name string
	if po.Name == "admin" {
		name = "Admin"
	}

	return &biz.Account{
		Id:    po.ID,
		Name:  name,
		Email: po.Email,
		Age:   po.Age,
		Sex:   po.Sex,
	}, nil
}

func (a *accountRepo) Login(ctx context.Context, username string) (*biz.Account, error) {
	po, err := a.data.db.Account.Query().Where(account.Name(username)).Only(ctx)
	if err != nil {
		// 没有则创建一个
		po, err = a.data.db.Account.Create().SetName(username).SetEmail("").SetSex(false).Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	return &biz.Account{
		Id:    po.ID,
		Name:  po.Name,
		Email: po.Email,
		Age:   po.Age,
		Sex:   po.Sex,
	}, nil
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
		return nil, err
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
