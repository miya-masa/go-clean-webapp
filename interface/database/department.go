package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/miya-masa/go-clean-webapp/domain/entity"
)

type departmentRepository struct {
	db *sqlx.Tx
}

func NewDepartment(db *sqlx.Tx) entity.DepartmentRepository {
	return &departmentRepository{db: db}
}

func (u *departmentRepository) Find(ctx context.Context, id string) (*entity.Department, error) {
	department := &entity.Department{}
	query := `
	select
		uuid,
		name
	from department
	where
		uuid = $1`
	if err := u.db.GetContext(ctx, department, query, id); err != nil {
		return nil, err
	}
	return department, nil
}

func (u *departmentRepository) Store(ctx context.Context, department *entity.Department) (*entity.Department, error) {
	if _, err := u.db.NamedExec("INSERT INTO department(uuid, name) VALUES(:uuid, :name)", department); err != nil {
		return nil, err
	}
	return department, nil
}

func (u *departmentRepository) Delete(ctx context.Context, id string) (int, error) {
	var affected int64
	r, err := u.db.Exec("DELETE FROM department where uuid = ?", id)
	if err != nil {
		return 0, err
	}
	i, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}
	affected += i
	return int(affected), nil
}
