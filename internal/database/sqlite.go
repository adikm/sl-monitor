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

type sqlite struct {
	dbName string
}

func NewSqlite(dbName string) Database {
	return &sqlite{dbName}
}

func (sqlite sqlite) Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./"+sqlite.dbName)
	if err != nil {
		log.Fatal("Unable to connect to db ", err)
	}
	return db
}

func (sqlite sqlite) Migrate(db *sql.DB) {
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{DatabaseName: sqlite.dbName})
	if err != nil {
		log.Fatal("Unable to connect to db ", err)
	}
	migrator, err := migrate.NewWithDatabaseInstance("file://assets/migrations", sqlite.dbName, instance)
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
