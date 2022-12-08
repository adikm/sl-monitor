package trafikverket

import (
	"encoding/json"
	"testing"
)

func Test_buildStationsRequest(t *testing.T) {
	t.Run("request login object correct", func(t *testing.T) {
		authKey := "testAuthKey"
		result := buildStationsRequest(authKey)

		want := login{authKey}
		if result.Login != want {
			t.Errorf("buildStationsRequest() login = %v, want %v", result.Login, want)
		}
	})
}

func TestRemoteService_FetchStations(t *testing.T) {
	service := APIService{&remoteClientStub{}}

	stations, err := service.FetchStations("")

	if err != nil {
		t.Fatal(err)
	}

	if len(stations) == 0 {
		t.Errorf("received slice length = %v, want %v", 0, 1)
	}

	expectedStation := Station{Name: "TestStation", Id: "TS"}
	if stations[0] != expectedStation {
		t.Errorf("received object = %v, want %v", stations[0], expectedStation)
	}
}

type remoteClientStub struct {
}

func (_ *remoteClientStub) post(r *request, dst interface{}) error {
	stubbedData := "{ \"RESPONSE\":{\"RESULT\":[{\"TrainStation\":[{\"AdvertisedLocationName\":\"TestStation\",\"LocationSignature\":\"TS\"}]}]}}"
	json.Unmarshal([]byte(stubbedData), &dst)
	return nil
}

var _ client = &remoteClientStub{}
