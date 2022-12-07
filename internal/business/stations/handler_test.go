package stations

import (
	"net/http"
	"net/http/httptest"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/config"
	"testing"
)

func TestStationHandler_FetchStations(t *testing.T) {
	req, err := http.NewRequest("GET", "/stations", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewHandler(&config.Config{}, &trafikverket.ServiceStub{}).FetchStations)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
