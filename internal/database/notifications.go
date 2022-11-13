package database

import (
	"context"
	"sl-monitor/internal"
	"time"
)

func (db *DB) CreateNotification(email string, timestamp time.Time, weekdays internal.WeekdaysSum) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "INSERT INTO notifications (email, timestamp, weekdays) VALUES ($1, $2, $3)"

	_, err := db.ExecContext(ctx, query, email, timestamp, weekdays)
	if err != nil {
		return err
	}
	return nil
}
