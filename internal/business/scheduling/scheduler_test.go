package scheduling

import (
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/business/users"
	"testing"
)

func TestScheduler_ScheduleNotifications(t *testing.T) {
	scheduler := Scheduler{nService: &notifications.NotificationServiceStub{}, uService: users.UserServiceStub{}, Mailer: nil}
	got := scheduler.ScheduleNotifications()

	if len(got) != 3 {
		t.Errorf("ScheduleNotifications() returned wrong length got = %v, want %v", len(got), 3)
	}
	for _, result := range got {
		if result.success != true {
			t.Errorf("ScheduleNotifications() returned success=false for notificationId=%v", result.notificationId)
		}
	}
}
