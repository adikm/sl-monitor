package stations

import (
	"net/http"
)

func Routes(sh *Handler) {

	http.HandleFunc("/stations", sh.FetchStations)

}
