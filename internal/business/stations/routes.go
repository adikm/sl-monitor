package stations

import (
	"net/http"
	request "sl-monitor/internal/server"
)

func Routes(sh *Handler) {

	http.HandleFunc("/stations", request.MustBe(http.MethodGet, sh.FetchStations))

}
