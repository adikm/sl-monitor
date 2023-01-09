package scheduling

import (
	"fmt"
	"sl-monitor/internal/business/notifications"
	"testing"
)

func TestScheduler_ScheduleNotifications(t *testing.T) {
	scheduler := Scheduler{Service: &notifications.NotificationServiceStub{}}
	got := scheduler.ScheduleNotifications()

	fmt.Printf("%v", got)
	if len(got) != 3 {
		t.Errorf("ScheduleNotifications() returned wrong length got = %v, want %v", len(got), 3)
	}
	for _, result := range got {
		if result.success != true {
			t.Errorf("ScheduleNotifications() returned success=false for notificationId=%v", result.notificationId)
		}
	}
}
