package usecase

import (
	"context"

	"github.com/miya-masa/go-clean-webapp/domain/entity"
	"github.com/miya-masa/go-clean-webapp/transaction"
)

type DepartmentStoreInput struct {
	UUID string
	Name string
}

type DepartmentInputPort interface {
	Find(ctx context.Context, id string) (*entity.Department, error)
	Store(ctx context.Context, in *DepartmentStoreInput) error
	Delete(ctx context.Context, id string) (int, error)
}

func NewDepartmentInteractor(dr entity.DepartmentRepository, tx transaction.Transaction) DepartmentInputPort {
	return &departmentInteractor{departmentRepository: dr, tx: tx}
}

type departmentInteractor struct {
	departmentRepository entity.DepartmentRepository
	tx                   transaction.Transaction
}

func (u *departmentInteractor) Find(ctx context.Context, id string) (*entity.Department, error) {
	return u.departmentRepository.Find(ctx, id)
}

func (u *departmentInteractor) Store(ctx context.Context, in *DepartmentStoreInput) error {
	_, err := u.tx.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		return nil, u.departmentRepository.Store(ctx, &entity.Department{
			UUID: genUUID(),
			Name: in.Name,
		})
	})
	return err
}

func (u *departmentInteractor) Delete(ctx context.Context, id string) (int, error) {
	return u.departmentRepository.Delete(ctx, id)
}
