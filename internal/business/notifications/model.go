package notifications

import (
	"sl-monitor/internal"
	"time"
)

type Store interface {
	Create(email string, timestamp time.Time, weekdays internal.WeekdaysSum) (int64, error)
}

type Notification struct {
	id        int64
	email     string
	timestamp time.Time
	weekdays  []internal.Weekday
}
