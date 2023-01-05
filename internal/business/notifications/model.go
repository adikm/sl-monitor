package notifications

import (
	"sl-monitor/internal"
	"time"
)

type Store interface {
	Create(timestamp time.Time, weekdays internal.WeekdaysSum, userId int) (int, error)
	FindByUserId(userId int) (*[]Notification, error)
	FindAll(dayOfWeek internal.Weekday) (*[]Notification, error)
}

type Service interface {
	Create(timestamp time.Time, weekdays internal.WeekdaysSum, userId int) (*Notification, error)
	FindAllForWeekday(weekday internal.Weekday) (*[]Notification, error)
	findByUserId(userId int) (*[]Notification, error)
}

type Notification struct {
	Id        int                `json:"id"`
	Timestamp time.Time          `json:"timestamp"`
	Weekdays  []internal.Weekday `json:"weekdays"`
	UserId    int                `json:"userId"`
}

var _ Store = &NotificationStore{}
