// +build wireinject

package main

// The build tag makes sure the stub is not built in the final build.

import (
	"context"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/miya-masa/go-clean-webapp/interface/database"
	"github.com/miya-masa/go-clean-webapp/interface/presenter"
	"github.com/miya-masa/go-clean-webapp/usecase"
	"github.com/miya-masa/go-clean-webapp/web"
)

func setupApplication(ctx context.Context) (Application, error) {
	wire.Build(applicationSet, web.NewAccountHandler, usecase.NewAccountInteractor, database.NewAccount, presenter.AccountPresenter{}, db)
	return Application{}, nil
}

func db(ctx context.Context) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", "user=miya password=miya dbname=miya sslmode=disable")
}
