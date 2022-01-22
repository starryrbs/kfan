package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Account struct {
	Id    int
	Name  string
	Email string
	Age   int
	Sex   bool
}

type AccountRepo interface {
	ListAccount(context.Context) ([]*Account, error)
	CreateAccount(ctx context.Context, account *Account) (*Account, error)
	Login(ctx context.Context, username string) (*Account, error)
	GetAccount(ctx context.Context, id int) (*Account, error)
}

type AccountUseCase struct {
	Repo AccountRepo
	log  *log.Helper
}

func NewAccountUseCase(repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{Repo: repo, log: log.NewHelper(log.With(logger, "module", "useCase/account"))}
}
