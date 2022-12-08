package database

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Database interface {
	Connect() (*sql.DB, error)
	Migrate(db *sql.DB) error
}

type sqlite struct {
	dbName string
}

func NewSqlite(dbName string) Database {
	return &sqlite{dbName}
}

func (sqlite sqlite) Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./"+sqlite.dbName)
	if err != nil {
		log.Print("Unable to connect to db")
		return nil, err
	}
	return db, nil
}

func (sqlite sqlite) Migrate(db *sql.DB) error {
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{DatabaseName: sqlite.dbName})
	if err != nil {
		log.Print("Unable to connect to db")
		return err
	}
	migrator, err := migrate.NewWithDatabaseInstance("file://assets/migrations", sqlite.dbName, instance)
	if err != nil {
		log.Print("Unable to prepare migrations")
		return err
	}

	err = migrator.Up()
	switch {
	case errors.Is(err, migrate.ErrNoChange):
		break
	case err != nil:
		log.Print("Unable to migrate")
		return err
	}
	return nil
}
