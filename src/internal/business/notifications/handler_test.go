package notifications

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sl-monitor/internal"
	"sl-monitor/internal/cache"
	"testing"
	"time"
)

func TestNotificationHandler_findForCurrentUser(t *testing.T) {
	cache.InitStub()

	req, err := http.NewRequest("GET", "/notifications/all", nil)
	req.AddCookie(&http.Cookie{Name: "session_token"})

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewHandler(&NotificationService{&NotificationStoreStub{}}).findForCurrentUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var got []Notification
	decoder := json.NewDecoder(rr.Body)
	decoder.Decode(&got)

	want := []Notification{{Id: 1, Timestamp: time.Unix(12345, 0), Weekdays: []internal.Weekday{internal.Monday, internal.Wednesday}, UserId: 0}}

	if want[0].Id != got[0].Id || !reflect.DeepEqual(want[0].Weekdays, got[0].Weekdays) || want[0].UserId != got[0].UserId {
		t.Errorf("findForCurrentUser() got = %v, want %v", got, want)
	}
}
