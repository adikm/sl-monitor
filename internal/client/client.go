package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sl-monitor/internal/config"
)

func post(r *request, result any) error {
	xmlRequestData, err := xml.Marshal(r)
	if err != nil {
		log.Println(err)
		return err
	}

	resp, err := http.Post(config.Cfg.TrafficAPI.URL, "text/xml", bytes.NewBuffer(xmlRequestData))
	if err != nil {
		log.Println(err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Status not 200. %s", string(body)))
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
