package notifications

import (
	"context"
	"database/sql"
	"log"
	"sl-monitor/internal"
	"time"
)

type NotificationStore struct {
	*sql.DB
}

func NewStore(db *sql.DB) *NotificationStore {
	return &NotificationStore{db}
}

func (h *NotificationStore) Create(timestamp time.Time, weekdays internal.WeekdaysSum, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "INSERT INTO notifications (timestamp, weekdays, user_id) VALUES ($1, $2, $3);"

	id, err := h.ExecContext(ctx, query, timestamp, weekdays, userId)
	if err != nil {
		return 0, err
	}
	insertId, err := id.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertId), nil
}

func (h *NotificationStore) FindByUserId(userId int) (*[]Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	query := "SELECT * FROM notifications WHERE user_id = $1"

	result, err := h.QueryContext(ctx, query, userId)

	var notifications []Notification

	for result.Next() {
		var n Notification
		var weekdays int
		if err := result.Scan(&n.Id, &n.Timestamp, &n.UserId, &weekdays); err != nil {
			log.Println(err.Error())
		}
		n.Weekdays = internal.WeekdaysSum(weekdays).AsWeekdays()
		notifications = append(notifications, n)
	}

	if err != nil {
		return nil, err
	}
	return &notifications, err
}
