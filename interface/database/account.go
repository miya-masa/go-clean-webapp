package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/miya-masa/go-clean-webapp/domain/entity"
)

type accountRepository struct {
	db *sqlx.DB
}

func NewAccount(db *sqlx.DB) entity.AccountRepository {
	return &accountRepository{db: db}
}

func (u *accountRepository) Find(ctx context.Context, id string) (*entity.Account, error) {
	account := &entity.Account{}
	query := `select
		a.uuid as uuid,
		a.first_name as first_name,
		a.last_name as last_name
	from account as a
	where
		a.uuid = $1`

	if err := u.db.GetContext(ctx, account, query, id); err != nil {
		return nil, err
	}
	return account, nil
}

func (u *accountRepository) Store(ctx context.Context, account *entity.Account) (*entity.Account, error) {

	var dao interface {
		NamedExec(query string, arg interface{}) (sql.Result, error)
	}
	dao, ok := GetTx(ctx)
	if !ok {
		dao = u.db
	}
	if _, err := dao.NamedExec("INSERT INTO account(uuid, department_uuid, first_name, last_name) VALUES(:uuid, :department.uuid, :first_name, :last_name)", account); err != nil {
		return nil, err
	}
	return account, nil
}

func (u *accountRepository) Delete(ctx context.Context, id string) (int, error) {
	tx, err := u.db.Beginx()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var affected int64
	r, err := tx.Exec("DELETE FROM account where uuid = $1", id)
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
