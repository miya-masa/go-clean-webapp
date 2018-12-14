package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/miya-masa/go-clean-webapp/interface/database"
	"github.com/miya-masa/go-clean-webapp/usecase"
	"github.com/miya-masa/go-clean-webapp/web"
	"github.com/urfave/cli"
)

var (
	Version  string
	Revision string
)
var logger = log.New(os.Stdout, "tx", log.LstdFlags)

func main() {

	app := cli.NewApp()
	app.Name = "gotxsand"
	app.Version = fmt.Sprintf("%v-%v", Version, Revision)
	app.Commands = []cli.Command{
		{
			Name:      "serve",
			ShortName: "s",
			Action: func(c *cli.Context) error {
				logger.Printf("start serve")
				return serve()
			},
		},
		{
			Name:      "init",
			ShortName: "i",
			Action: func(c *cli.Context) error {
				logger.Printf("initialize data")
				return initializeData()
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func serve() error {

	db, err := sqlx.Connect("postgres", "user=miya password=miya dbname=miya sslmode=disable")
	if err != nil {
		return err
	}

	accountRepository := database.NewAccount(db)

	uh := &web.AccountHandler{
		Usecase: usecase.NewAccountInteractor(accountRepository, database.NewTransaction(db)),
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", uh.Post)
		r.Get("/{accountUUID}", uh.Get)
		r.Delete("/{accountUUID}", uh.Delete)
	})

	http.ListenAndServe(":8080", r)
	return nil
}
