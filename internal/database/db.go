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

type DB struct {
	*sql.DB
}

func Connect(dbName string) *DB {
	db, err := sql.Open("sqlite3", "./"+dbName)
	if err != nil {
		log.Fatal("Unable to connect to db ", err)
	}
	return &DB{db}
}

func Migrate(db *DB, dbName string) {
	instance, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{DatabaseName: dbName})
	if err != nil {
		log.Fatal("Unable to connect to db ", err)
	}
	migrator, err := migrate.NewWithDatabaseInstance("file://assets/migrations", dbName, instance)
	if err != nil {
		log.Fatal("Unable to prepare migrations ", err)
	}

	err = migrator.Up()
	switch {
	case errors.Is(err, migrate.ErrNoChange):
		break
	case err != nil:
		log.Fatal("Unable to migrate ", err)
	}
}
