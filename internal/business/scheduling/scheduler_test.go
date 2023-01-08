package scheduling

import (
	"sl-monitor/internal/business/notifications"
	"testing"
)

func TestScheduler_DoIt(t *testing.T) {
	scheduler := Scheduler{Service: &notifications.NotificationServiceStub{}}
	scheduler.ScheduleNotifications()
}
