package notifications

import (
	"sl-monitor/internal"
	"time"
)

type Store interface {
	Create(email string, timestamp time.Time, weekdays internal.WeekdaysSum) (int, error)
}

type Notification struct {
	id        int
	email     string
	timestamp time.Time
	weekdays  []internal.Weekday
}
