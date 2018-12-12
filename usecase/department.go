package usecase

import (
	"context"

	"github.com/miya-masa/go-clean-webapp/domain/entity"
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

func NewDepartmentInteractor(dr entity.DepartmentRepository) DepartmentInputPort {
	return &departmentInteractor{departmentRepository: dr}
}

type departmentInteractor struct {
	departmentRepository entity.DepartmentRepository
}

func (u *departmentInteractor) Find(ctx context.Context, id string) (*entity.Department, error) {
	return u.departmentRepository.Find(ctx, id)
}

func (u *departmentInteractor) Store(ctx context.Context, in *DepartmentStoreInput) error {
	return u.departmentRepository.Store(ctx, &entity.Department{
		UUID: genUUID(),
		Name: in.Name,
	})
}

func (u *departmentInteractor) Delete(ctx context.Context, id string) (int, error) {
	return u.departmentRepository.Delete(ctx, id)
}
