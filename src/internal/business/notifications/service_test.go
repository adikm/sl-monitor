package notifications

import (
	"reflect"
	"sl-monitor/internal"
	"testing"
	"time"
)

func TestNotificationService_Create(t *testing.T) {
	s := NotificationService{&NotificationStoreStub{}}

	want := &Notification{
		Id:          125,
		Timestamp:   time.Unix(12346, 0),
		Weekdays:    []internal.Weekday{internal.Saturday, internal.Sunday},
		UserId:      54,
		StationCode: "Hnd",
	}

	got, _ := s.Create(time.Unix(12346, 0), 96, 54, "Hnd")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindAllForWeekday() got = %v, want %v", got, want)
	}
}

func TestNotificationService_FindAllForWeekday(t *testing.T) {
	s := NotificationService{&NotificationStoreStub{}}

	want := &[]Notification{{
		Id:        2,
		Timestamp: time.Unix(12346, 0),
		Weekdays:  []internal.Weekday{internal.Wednesday},
		UserId:    1,
	}}

	got, _ := s.FindAllForWeekday(internal.Wednesday)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindAllForWeekday() got = %v, want %v", got, want)
	}
}

func TestNotificationService_findByUserId(t *testing.T) {
	s := NotificationService{&NotificationStoreStub{}}

	want := &[]Notification{{
		Id:        1,
		Timestamp: time.Unix(12345, 0),
		Weekdays:  []internal.Weekday{internal.Monday, internal.Wednesday},
		UserId:    123,
	}}

	got, _ := s.findByUserId(123)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindAllForWeekday() got = %v, want %v", got, want)
	}
}
