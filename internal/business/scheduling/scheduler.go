package scheduling

import (
	"fmt"
	"log"
	"sl-monitor/internal"
	"sl-monitor/internal/business/notifications"
	"sync"
	"time"
)

type Scheduler struct {
	Service notifications.Service
}

type Result struct {
	success        bool
	notificationId int
}

// ScheduleNotifications fanin fanout pattern.
// this method is blocking
func (s *Scheduler) ScheduleNotifications() []Result {
	today := s.today()
	notificationsToSchedule, err := s.Service.FindAllForWeekday(today)
	if err != nil {
		log.Println(err)
		return nil
	}

	var results []Result
	var wg sync.WaitGroup
	for _, n := range *notificationsToSchedule {
		now := time.Now()
		executionDate := time.Date(now.Year(), now.Month(), now.Day(), n.Timestamp.Hour(), n.Timestamp.Minute(), n.Timestamp.Second(), 0, time.Local) // always today and time that's saved in db
		resultChan := make(chan Result, 1)
		wg.Add(1)
		go s.scheduleNotification(n, executionDate, &wg, resultChan)
		wg.Add(1)
		go s.collectResult(&results, resultChan, &wg)
	}
	wg.Wait()
	fmt.Printf("Finished with results %v \r\n", results)

	return results
}

func (s *Scheduler) scheduleNotification(n notifications.Notification, until time.Time, wg *sync.WaitGroup, sampleChan chan Result) {
	defer wg.Done()

	defer close(sampleChan)
	var w sync.WaitGroup
	w.Add(1)
	fmt.Printf("Scheduled notification id=%d on %s \r\n", n.Id, until)
	time.Sleep(time.Until(until))
	go s.sendNotification(n, sampleChan, &w)
	w.Wait()
}

func (s *Scheduler) collectResult(results *[]Result, resultChan chan Result, wg *sync.WaitGroup) { // consumer
	defer wg.Done()
	for s := range resultChan {
		*results = append(*results, s)
	}
}

func (s *Scheduler) sendNotification(n notifications.Notification, channel chan<- Result, wg *sync.WaitGroup) { // producer
	defer wg.Done()
	fmt.Printf("Sending notification id =%d \r\n", n.Id)
	fmt.Printf("EMAILED!! %v \r\n", n) // email it
	channel <- Result{success: true, notificationId: n.Id}
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
