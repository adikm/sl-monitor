package scheduling

import (
	"fmt"
	"sl-monitor/internal"
	"sl-monitor/internal/business/notifications"
	"sync"
	"time"
)

type Scheduler struct {
	Service notifications.Service
}

type result struct {
	success bool
	no      notifications.Notification
}

// ScheduleNotifications fanin fanout pattern
func (s *Scheduler) ScheduleNotifications() {
	today := s.today()
	notificationsToSchedule, err := s.Service.FindAllForWeekday(today)
	if err != nil {

	}

	var results []result
	var wg sync.WaitGroup
	for _, n := range *notificationsToSchedule {
		now := time.Now()
		executionDate := time.Date(now.Year(), now.Month(), now.Day(), n.Timestamp.Hour(), n.Timestamp.Minute(), n.Timestamp.Second(), 0, time.Local) // always today and time that's saved in db
		resultChan := make(chan result, 1)
		wg.Add(1)
		go s.scheduleNotification(n, executionDate, &wg, resultChan)
		wg.Add(1)
		go s.collectResult(&results, resultChan, &wg)
	}
	wg.Wait()
	fmt.Println(results)

	fmt.Println("finito")
}

func (s *Scheduler) scheduleNotification(n notifications.Notification, until time.Time, wg *sync.WaitGroup, sampleChan chan result) {
	defer wg.Done()

	defer close(sampleChan)
	var w sync.WaitGroup
	w.Add(1)
	fmt.Printf("starting and waiting until %s \r\n", until)
	time.Sleep(time.Until(until))
	go s.sendNotification(n, sampleChan, &w)
	w.Wait()
}

func (s *Scheduler) collectResult(results *[]result, resultChan chan result, wg *sync.WaitGroup) { // consumer
	defer wg.Done()
	for s := range resultChan {
		*results = append(*results, s)
	}
}

func (s *Scheduler) sendNotification(n notifications.Notification, channel chan<- result, wg *sync.WaitGroup) { // producer
	defer wg.Done()
	fmt.Printf("EMAILED!! %v \r\n", n) // email it
	channel <- result{success: true, no: n}
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
