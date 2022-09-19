package main

import (
	"database/sql"
	"log"

	config "github.com/abekoh/go-saas-ovation"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cnf := config.Load()
	db, err := sql.Open("postgres", cnf.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file:///migration", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
