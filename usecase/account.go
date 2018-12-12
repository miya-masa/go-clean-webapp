package usecase

import (
	"context"

	"github.com/miya-masa/go-clean-webapp/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type AccountStoreInput struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DepartmentUUID string `json:"department_uuid"`
}

type AccountInputPort interface {
	Find(ctx context.Context, id string) (*entity.Account, error)
	Store(ctx context.Context, in *AccountStoreInput) (*entity.Account, error)
	Delete(ctx context.Context, id string) (int, error)
}

func NewAccountInteractor(ar entity.AccountRepository, dr entity.DepartmentRepository) AccountInputPort {
	return &accountInteractor{
		accountRepository:    ar,
		departmentRepository: dr,
	}
}

type accountInteractor struct {
	accountRepository    entity.AccountRepository
	departmentRepository entity.DepartmentRepository
}

func (u *accountInteractor) Find(ctx context.Context, id string) (*entity.Account, error) {
	return u.accountRepository.Find(ctx, id)
}

func (u *accountInteractor) Store(ctx context.Context, in *AccountStoreInput) (*entity.Account, error) {
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
}

func (u *accountInteractor) Delete(ctx context.Context, id string) (int, error) {
	return u.accountRepository.Delete(ctx, id)
}

func genUUID() string {
	return uuid.NewV4().String()
}
