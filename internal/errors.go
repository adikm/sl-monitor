package internal

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"sl-monitor/internal/server/response"
	"strings"
)

type JsonCommon struct {
	logger *log.Logger
}

func NewJsonCommon(logger *log.Logger) *JsonCommon {
	return &JsonCommon{logger}
}

func (common *JsonCommon) ErrorMessage(w http.ResponseWriter, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		common.logger.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (common *JsonCommon) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	common.logger.Print(trace)

	message := "The server encountered a problem and could not process your request"
	common.ErrorMessage(w, http.StatusInternalServerError, message, nil)
}

func (common *JsonCommon) NotFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	common.ErrorMessage(w, http.StatusNotFound, message, nil)
}

func (common *JsonCommon) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	common.ErrorMessage(w, http.StatusMethodNotAllowed, message, nil)
}

func (common *JsonCommon) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	common.ErrorMessage(w, http.StatusBadRequest, err.Error(), nil)
}
