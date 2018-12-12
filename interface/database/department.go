package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/miya-masa/go-clean-webapp/domain/entity"
)

type departmentRepository struct {
	db *sqlx.DB
}

func NewDepartment(db *sqlx.DB) entity.DepartmentRepository {
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

func (u *departmentRepository) Store(ctx context.Context, department *entity.Department) error {
	tx, err := u.db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.NamedExec("INSERT INTO department(uuid, name) VALUES(:uuid, :name)", department); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (u *departmentRepository) Delete(ctx context.Context, id string) (int, error) {
	tx, err := u.db.Beginx()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var affected int64
	r, err := tx.Exec("DELETE FROM department where uuid = ?", id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	i, err := r.RowsAffected()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	affected += i
	tx.Commit()
	return int(affected), nil
}
