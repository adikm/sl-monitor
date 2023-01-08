package notifications

import (
	"sl-monitor/internal"
	"time"
)

type NotificationServiceStub struct {
}

func (s *NotificationServiceStub) Create(timestamp time.Time, weekdays internal.WeekdaysSum, userId int) (*Notification, error) {
	return nil, nil
}

func (s *NotificationServiceStub) findByUserId(userId int) (*[]Notification, error) {
	return nil, nil
}

func (s *NotificationServiceStub) FindAllForWeekday(weekday internal.Weekday) (*[]Notification, error) {
	now := time.Now()
	var notifications []Notification
	notifications = append(notifications, Notification{Id: 5, Timestamp: now.Add(time.Second * time.Duration(5)), Weekdays: []internal.Weekday{internal.Sunday}, UserId: 0})
	notifications = append(notifications, Notification{Id: 3, Timestamp: now.Add(time.Second * time.Duration(3)), Weekdays: []internal.Weekday{internal.Sunday}, UserId: 0})
	notifications = append(notifications, Notification{Id: 7, Timestamp: now.Add(time.Second * time.Duration(7)), Weekdays: []internal.Weekday{internal.Sunday}, UserId: 0})
	return &notifications, nil
}
