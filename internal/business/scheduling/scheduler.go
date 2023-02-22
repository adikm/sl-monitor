package scheduling

import (
	"log"
	"sl-monitor/internal"
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/smtp"
	"sync"
	"time"
)

type Scheduler struct {
	nService notifications.Service
	sender   *Sender
	mailer   *smtp.Mailer
}

func NewScheduler(service notifications.Service, sender *Sender, mailer *smtp.Mailer) *Scheduler {
	return &Scheduler{service, sender, mailer}
}

type Result struct {
	success        bool
	notificationId int
}

// ScheduleNotifications fanin fanout pattern.
// this method is blocking
func (s *Scheduler) ScheduleNotifications() []Result {
	today := internal.TodayWeekday()
	notificationsToSchedule, err := s.nService.FindAllForWeekday(today)
	if err != nil {
		log.Println(err)
		return nil
	}

	var results []Result
	var wg sync.WaitGroup
	for _, n := range *notificationsToSchedule {
		now := time.Now()
		executionDate := s.getExecutionDate(now, n) // always today and time that's saved in db
		if time.Now().After(executionDate) {        // if current time is already after the expected schedule time, skip - helpful during redeployment
			log.Printf("Ignoring scheduling notificationId=%d because the time is after the expected schedule date\n", n.Id)
			continue
		}
		resultChan := make(chan Result, 1)
		wg.Add(1)
		go s.scheduleNotification(n, executionDate, &wg, resultChan)
		wg.Add(1)
		go s.collectResult(&results, resultChan, &wg)
	}
	wg.Wait()
	log.Printf("Finished with results %v \r\n", results)

	return results
}

func (s *Scheduler) getExecutionDate(now time.Time, n notifications.Notification) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), n.Timestamp.Hour(), n.Timestamp.Minute(), n.Timestamp.Second(), 0, time.Local)
}

func (s *Scheduler) scheduleNotification(n notifications.Notification, until time.Time, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done()

	defer close(resultChan)
	var w sync.WaitGroup
	w.Add(1)
	log.Printf("Scheduled notification id=%d on %s \r\n", n.Id, until)
	time.Sleep(time.Until(until))
	go s.performScheduledNotification(n, resultChan, &w)
	w.Wait()
}

func (s *Scheduler) collectResult(results *[]Result, resultChan chan Result, wg *sync.WaitGroup) { // consumer
	defer wg.Done()
	for s := range resultChan {
		*results = append(*results, s)
	}
}

func (s *Scheduler) performScheduledNotification(n notifications.Notification, channel chan<- Result, wg *sync.WaitGroup) { // producer
	defer wg.Done()
	log.Printf("Sending notification id =%d \r\n", n.Id)

	to, body := s.sender.prepareNotificationMail(n.UserId, n.StationCode)
	if body == nil {
		channel <- Result{success: false, notificationId: n.Id}
		return
	}

	if s.mailer != nil {
		s.mailer.SendMail(to, *body)
	}
	log.Printf("EMAILED!! %v \r\n", n) // email it
	channel <- Result{success: true, notificationId: n.Id}
}
