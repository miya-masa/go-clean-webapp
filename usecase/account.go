package usecase

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/miya-masa/go-clean-webapp/domain/entity"
	"github.com/miya-masa/go-clean-webapp/interface/database"
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

type txAccountInteractor struct {
	db *sqlx.DB
}

func NewAccountInteractorTx(db *sqlx.DB) AccountInputPort {
	return &txAccountInteractor{db: db}
}

func (u *txAccountInteractor) Store(ctx context.Context, in *AccountStoreInput) (*entity.Account, error) {
	v, err := database.DoInTx(u.db, func(tx *sqlx.Tx) (interface{}, error) {
		ar := database.NewAccount(tx)
		return NewAccountInteractor(ar).Store(ctx, in)
	})
	return v.(*entity.Account), err
}

func (u *txAccountInteractor) Find(ctx context.Context, id string) (*entity.Account, error) {
	v, err := database.DoInTx(u.db, func(tx *sqlx.Tx) (interface{}, error) {
		ar := database.NewAccount(tx)
		return NewAccountInteractor(ar).Find(ctx, id)
	})
	return v.(*entity.Account), err
}

func (u *txAccountInteractor) Delete(ctx context.Context, id string) (int, error) {
	num, err := database.DoInTx(u.db, func(tx *sqlx.Tx) (interface{}, error) {
		ar := database.NewAccount(tx)
		return NewAccountInteractor(ar).Delete(ctx, id)
	})
	return num.(int), err
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
