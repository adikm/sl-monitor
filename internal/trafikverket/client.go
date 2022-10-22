package trafikverket

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func post(r *request, dst interface{}) error {
	requestBody, err := xml.Marshal(r)
	if err != nil {
		return fmt.Errorf("couldn't prepare request body: %w", err)
	}

	resp, err := http.Post("https://api.trafikinfo.trafikverket.se/v2/data.json", "text/xml", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("trafikverket POST request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
