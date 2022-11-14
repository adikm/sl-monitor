package response

import (
	"fmt"
	"net/http"
	"runtime/debug"
	customlogger "sl-monitor/internal/logger"
	"strings"
)

func ErrorMessage(w http.ResponseWriter, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		customlogger.GetInstance().Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	customlogger.GetInstance().Print(trace)

	message := "The server encountered a problem and could not process your request"
	ErrorMessage(w, http.StatusInternalServerError, message, nil)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	ErrorMessage(w, http.StatusNotFound, message, nil)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	ErrorMessage(w, http.StatusMethodNotAllowed, message, nil)
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	ErrorMessage(w, http.StatusBadRequest, err.Error(), nil)
}
