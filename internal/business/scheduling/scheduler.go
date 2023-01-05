package scheduling

import (
	"fmt"
	"sl-monitor/internal"
	"sl-monitor/internal/business/notifications"
	"time"
)

type Scheduler struct {
	Service notifications.Service
}

// DoIt TODO Should be run every midnight
func (s *Scheduler) DoIt() {
	today := s.today()
	result, err := s.Service.FindAllForWeekday(today)
	if err != nil {

	}
	for _, n := range *result {
		now := time.Now()
		executionDate := time.Date(now.Year(), now.Month(), now.Day(), n.Timestamp.Hour(), n.Timestamp.Minute(), 0, 0, time.Local)
		go s.notify(n, executionDate)
	}
}

func (s *Scheduler) notify(n notifications.Notification, executionDate time.Time) {
	time.Sleep(time.Until(executionDate))
	fmt.Println(n) // DoIt TODO send notification
}

func (s *Scheduler) today() internal.Weekday {
	switch weekday := time.Now().Weekday(); weekday {
	case time.Monday:
		return internal.Monday
	case time.Tuesday:
		return internal.Tuesday
	case time.Wednesday:
		return internal.Wednesday
	case time.Thursday:
		return internal.Thursday
	case time.Friday:
		return internal.Friday
	case time.Saturday:
		return internal.Saturday
	case time.Sunday:
		return internal.Sunday
	default:
		return internal.Monday
	}
}
