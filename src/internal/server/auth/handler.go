package auth

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"sl-monitor/internal/business/users"
	"sl-monitor/internal/cache"
	"sl-monitor/internal/config"
	"sl-monitor/internal/server"
	"sl-monitor/internal/server/response"
	"strconv"
	"time"
)

type Handler struct {
	config *config.Config
	users  users.Service
}

func NewHandler(config *config.Config, users users.Service) *Handler {
	return &Handler{config, users}
}

func (ah *Handler) login(w http.ResponseWriter, r *http.Request) {
	type Credentials struct {
		Pwd   string `json:"password"`
		Email string `json:"email"`
	}

	var creds Credentials

	err := request.DecodeJSON(r, &creds)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	idPass, err := ah.users.FindPasswordByEmail(creds.Email)

	if err != nil {
		log.Println(err)
		response.ServerError(w, r, err)
		return
	}
	if idPass == nil || !users.CheckPasswordHash(creds.Pwd, idPass.Pwd) {
		response.ErrorMessage(w, http.StatusUnauthorized, "User or password incorrect", nil)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)
	cache.Instance.SetValue(sessionToken, strconv.Itoa(idPass.Id), 120*time.Second)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
}

func (ah *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			response.Unauthorized(w, r)
			return
		}
		response.BadRequest(w, r, err)
		return
	}
	sessionToken := cookie.Value

	cache.Instance.DeleteValue(sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}
