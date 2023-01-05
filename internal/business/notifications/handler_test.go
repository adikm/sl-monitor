package notifications

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sl-monitor/internal"
	"testing"
	"time"
)

func TestNotificationHandler_findForCurrentUser(t *testing.T) {
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

	var result []Notification
	decoder := json.NewDecoder(rr.Body)
	decoder.Decode(&result)

	if len(result) != 1 {
		t.Errorf("handler returned wrong value: got size %v want %v",
			len(result), 1)
	}

	if result[0].Id != 1 {
		t.Errorf("handler returned wrong value: got id %v want %v",
			result[0].Id, 1)
	}

	expectedTimestamp := time.Unix(12345, 0)
	if result[0].Timestamp != expectedTimestamp {
		t.Errorf("handler returned wrong value: got time %v want %v",
			result[0].Timestamp, expectedTimestamp)
	}

	if result[0].Weekdays[0] != internal.Monday || result[0].Weekdays[1] != internal.Wednesday {
		t.Errorf("handler returned wrong value: got weekdays %v want %v %v",
			result[0].Weekdays, internal.Monday, internal.Wednesday)
	}

	if result[0].UserId != 0 {
		t.Errorf("handler returned wrong value: got userId %v want %v",
			result[0].Id, 0)
	}
}
