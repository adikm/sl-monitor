package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"log"
	"sl-monitor/internal/config"
)

type Database interface {
	Connect() (*sql.DB, error)
	Migrate(db *sql.DB) error
}

type postgre struct {
	dbName   string
	host     string
	port     int
	user     string
	password string
}

func NewPostgre(db config.Database) Database {
	return &postgre{dbName: db.Name, host: db.Host, port: db.Port, user: db.User, password: db.Password}
}

func (p postgre) Connect() (*sql.DB, error) {
	sqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.user, p.password, p.dbName)

	db, err := sql.Open("postgres", sqlconn)
	if err != nil {
		log.Print("Unable to connect to db")
		return nil, err
	}
	return db, nil
}

func (p postgre) Migrate(db *sql.DB) error {

	instance, err := postgres.WithInstance(db, &postgres.Config{DatabaseName: p.dbName})
	if err != nil {
		return fmt.Errorf("unable to connect to db during migration: %s", err)
	}
	migrator, err := migrate.NewWithDatabaseInstance("file://assets/migrations", p.dbName, instance)
	migrator.Log = &directLogger{}

	if err != nil {
		return fmt.Errorf("unable to prepare migration: %s", err)
	}

	err = migrator.Up()
	switch {
	case errors.Is(err, migrate.ErrNoChange):
		break
	case err != nil:
		return fmt.Errorf("unable to migrate: %s", err)
	}
	return nil
}

type directLogger struct{}

func (l *directLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *directLogger) Verbose() bool {
	return false
}

var _ Database = &postgre{}
