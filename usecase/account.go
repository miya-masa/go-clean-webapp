package usecase

import (
	"context"

	"github.com/miya-masa/go-clean-webapp/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type AccountStoreInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AccountInputPort interface {
	Find(ctx context.Context, id string) (*entity.Account, error)
	Store(ctx context.Context, in *AccountStoreInput) (*entity.Account, error)
	Delete(ctx context.Context, id string) (int, error)
}

func NewAccountInteractor(ar entity.AccountRepository) AccountInputPort {
	return &accountInteractor{
		accountRepository: ar,
	}
}

type accountInteractor struct {
	accountRepository entity.AccountRepository
}

func (u *accountInteractor) Find(ctx context.Context, id string) (*entity.Account, error) {
	return u.accountRepository.Find(ctx, id)
}

func (u *accountInteractor) Store(ctx context.Context, in *AccountStoreInput) (*entity.Account, error) {
	return u.accountRepository.Store(ctx, entity.New(in.FirstName, in.FirstName))
}

func (u *accountInteractor) Delete(ctx context.Context, id string) (int, error) {
	return u.accountRepository.Delete(ctx, id)
}

func genUUID() string {
	return uuid.NewV4().String()
}
