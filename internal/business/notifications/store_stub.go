package notifications

import (
	"sl-monitor/internal"
	"time"
)

type NotificationStoreStub struct {
}

func (h *NotificationStoreStub) Create(timestamp time.Time, weekdays internal.WeekdaysSum, userId int) (int, error) {
	return 0, nil
}

func (h *NotificationStoreStub) FindByUserId(userId int) (*[]Notification, error) {
	return &[]Notification{{
		Id:        1,
		Timestamp: time.Unix(12345, 0),
		Weekdays:  []internal.Weekday{internal.Monday, internal.Wednesday},
		UserId:    userId,
	}}, nil
}

func (h *NotificationStoreStub) FindAll(weekday internal.Weekday) (*[]Notification, error) {
	return &[]Notification{}, nil
}

var _ Store = &NotificationStoreStub{}
