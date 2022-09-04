package db

import (
	"database/sql"
	"embed"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed schema/*.sql
var migrations embed.FS

func MigrateUp() {
	var db *sql.DB

	// setup database
	db, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	goose.SetBaseFS(migrations)

	if err := goose.Up(db, "schema"); err != nil {
		panic(err)
	}

	// run app
}
