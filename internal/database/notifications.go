package database

import (
	"context"
	"time"
)

func (db *DB) CreateNotification(email string, timestamp time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "INSERT INTO notifications (email, timestamp)	VALUES ($1, $2)"

	_, err := db.ExecContext(ctx, query, email, timestamp)
	if err != nil {
		return err
	}
	return nil
}
