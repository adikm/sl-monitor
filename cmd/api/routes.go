package main

import (
	"net/http"
	"sl-monitor/internal/server/response"
)

func DefaultRoutes() {
	http.HandleFunc("/", response.NotFound)
}
