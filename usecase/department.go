package usecase

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/miya-masa/go-clean-webapp/domain/entity"
	"github.com/miya-masa/go-clean-webapp/interface/database"
)

type DepartmentStoreInput struct {
	UUID string
	Name string
}

type DepartmentInputPort interface {
	Find(ctx context.Context, id string) (*entity.Department, error)
	Store(ctx context.Context, in *DepartmentStoreInput) (*entity.Department, error)
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

func (u *departmentInteractor) Store(ctx context.Context, in *DepartmentStoreInput) (*entity.Department, error) {
	return u.departmentRepository.Store(ctx, &entity.Department{
		UUID: genUUID(),
		Name: in.Name,
	})
}

func (u *departmentInteractor) Delete(ctx context.Context, id string) (int, error) {
	return u.departmentRepository.Delete(ctx, id)
}

type txDepartmentInteractor struct {
	db *sqlx.DB
}

func NewDepartmentInteractorTx(db *sqlx.DB) DepartmentInputPort {
	return &txDepartmentInteractor{db: db}
}

func (u *txDepartmentInteractor) Store(ctx context.Context, in *DepartmentStoreInput) (*entity.Department, error) {
	v, err := database.DoInTx(u.db, func(tx *sqlx.Tx) (interface{}, error) {
		dr := database.NewDepartment(tx)
		return NewDepartmentInteractor(dr).Store(ctx, in)
	})
	return v.(*entity.Department), err
}

func (u *txDepartmentInteractor) Find(ctx context.Context, id string) (*entity.Department, error) {
	v, err := database.DoInTx(u.db, func(tx *sqlx.Tx) (interface{}, error) {
		dr := database.NewDepartment(tx)
		return NewDepartmentInteractor(dr).Find(ctx, id)
	})
	return v.(*entity.Department), err
}

func (u *txDepartmentInteractor) Delete(ctx context.Context, id string) (int, error) {
	num, err := database.DoInTx(u.db, func(tx *sqlx.Tx) (interface{}, error) {
		dr := database.NewDepartment(tx)
		return NewDepartmentInteractor(dr).Delete(ctx, id)
	})
	return num.(int), err
}
