package trafikverket

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type client interface {
	post(r *request, dst interface{}) error
}

type remoteClient struct {
}

func (_ *remoteClient) post(r *request, dst interface{}) error {
	requestBody, err := xml.Marshal(r)
	if err != nil {
		return fmt.Errorf("couldn't prepare request body: %w", err)
	}

	resp, err := http.Post("https://api.trafikinfo.trafikverket.se/v2/data.json", "text/xml", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("trafikverket POST request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("trafikverket POST request failed when reading body: %w", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response from Trafficverket. Status code %d: %s", resp.StatusCode, string(body))
	}
	err = json.Unmarshal(body, &dst)
	if err != nil {
		return fmt.Errorf("trafikverket POST request failed when unmarshalling body: %w", err)
	}
	return nil
}

var _ client = &remoteClient{}
