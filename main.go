package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

func newApplication(ctx context.Context, ah *web.AccountHandler) (Application, error) {
	return Application{AccountHandler: ah}, nil
}

func serve() error {

	ctx := context.Background()
	app, err := setupApplication(ctx)
	if err != nil {
		return err
	}
	return app.Start()
}

type Application struct {
	AccountHandler *web.AccountHandler
}

func (a *Application) Start() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/accounts", func(r chi.Router) {
		r.Get("/", a.AccountHandler.List)
		r.Post("/", a.AccountHandler.Post)
		r.Get("/{accountUUID}", a.AccountHandler.Get)
		r.Delete("/{accountUUID}", a.AccountHandler.Delete)
	})

	http.ListenAndServe(":8080", r)
	return nil
}
