package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS "account" (
	"uuid" varchar(36),
	"department_uuid" varchar(36),
	"first_name" text,
	"last_name" text,
	PRIMARY KEY ("uuid"));
`

func initializeData() error {
	db, err := sqlx.Connect("postgres", "user=miya password=miya dbname=miya sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)
	return nil
}
