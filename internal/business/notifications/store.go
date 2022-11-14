package notifications

import (
	"context"
	"database/sql"
	"sl-monitor/internal"
	"time"
)

type NotificationStore struct {
	*sql.DB
}

func NewStore(db *sql.DB) *NotificationStore {
	return &NotificationStore{db}
}

func (h *NotificationStore) Create(email string, timestamp time.Time, weekdays internal.WeekdaysSum) (Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "INSERT INTO notifications (email, timestamp, weekdays) VALUES ($1, $2, $3);"

	id, err := h.ExecContext(ctx, query, email, timestamp, weekdays)
	if err != nil {
		return Notification{}, err
	}
	insertId, err := id.LastInsertId()
	if err != nil {
		return Notification{}, err
	}
	return Notification{insertId}, nil
}
