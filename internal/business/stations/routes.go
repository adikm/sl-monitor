package stations

import (
	"net/http"
)

func Routes(sh *StationHandler) {

	http.HandleFunc("/stations", sh.FetchStations)

}
