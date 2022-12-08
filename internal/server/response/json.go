package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data any) error {
	return JSONWithHeaders(w, status, data, nil)
}

func JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	return encoder.Encode(data)

}
