package notifications

import (
	"sl-monitor/internal"
	"time"
)

type NotificationService struct {
	store Store
}

func NewService(store Store) *NotificationService {
	return &NotificationService{store}
}

func (s *NotificationService) Create(timestamp time.Time, weekdays internal.WeekdaysSum, userId int) (*Notification, error) {
	normalizedTime := time.Date(1970, 1, 1, timestamp.Hour(), timestamp.Minute(), timestamp.Second(), 0, time.Local)
	id, err := s.store.Create(normalizedTime, weekdays, userId)

	if err != nil {
		return nil, err
	}

	return &Notification{id, timestamp, weekdays.AsWeekdays(), userId}, nil
}

func (s *NotificationService) FindAllForWeekday(weekday internal.Weekday) (*[]Notification, error) {
	return s.store.FindAll(weekday)
}

func (s *NotificationService) findByUserId(userId int) (*[]Notification, error) {
	return s.store.FindByUserId(userId)
}
