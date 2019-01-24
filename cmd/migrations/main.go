package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/kostozyb/orm-bench/internal/config"
	"github.com/labstack/gommon/log"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func applyMigrations(driver, cs, dbName string) error {
	db, err := sql.Open(driver, cs)
	if err != nil {
		return err
	}

	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./misc/migrations",
		dbName, dbDriver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return err
	}

	if err := db.Close(); err != nil {
		return err
	}

	return nil
}

func main() {
	c := config.Config{}

	if err := applyMigrations(c.GetDriver(), c.GetConnectionString(), c.GetDB()); err != nil {
		log.Panic(err)
	}

	log.Info("Migrations complete")
}
