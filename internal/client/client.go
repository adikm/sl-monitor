package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"sl-monitor/internal/config"
)

func post(r *request, result any) {
	xmlRequestData, err := xml.Marshal(r)
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.Post(config.Cfg.TrafficAPI.URL, "text/xml", bytes.NewBuffer(xmlRequestData))
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		log.Println("Status not 200", string(body))
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
		return
	}
}
