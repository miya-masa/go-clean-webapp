package usecase

import (
	"context"

	"github.com/miya-masa/go-clean-webapp/domain/entity"
	"github.com/miya-masa/go-clean-webapp/transaction"
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

func NewAccountInteractor(ar entity.AccountRepository, tx transaction.Transaction) AccountInputPort {
	return &accountInteractor{
		accountRepository: ar,
		trancaction:       tx,
	}
}

type accountInteractor struct {
	accountRepository entity.AccountRepository
	trancaction       transaction.Transaction
}

func (u *accountInteractor) Find(ctx context.Context, id string) (*entity.Account, error) {
	return u.accountRepository.Find(ctx, id)
}

func (u *accountInteractor) Store(ctx context.Context, in *AccountStoreInput) (*entity.Account, error) {
	v, err := u.trancaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		dep, err := u.departmentRepository.Find(ctx, in.DepartmentUUID)
		if err != nil {
			return nil, err
		}
		return u.accountRepository.Store(ctx, &entity.Account{
			UUID:       genUUID(),
			Department: dep,
			FirstName:  in.FirstName,
			LastName:   in.LastName,
		})
	})
	return v.(*entity.Account), err
}

func (u *accountInteractor) Delete(ctx context.Context, id string) (int, error) {
	return u.accountRepository.Delete(ctx, id)
}

func genUUID() string {
	return uuid.NewV4().String()
}
