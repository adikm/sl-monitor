package auth

import (
	"github.com/google/uuid"
	"net/http"
	"sl-monitor/internal/business/users"
	"sl-monitor/internal/config"
	"sl-monitor/internal/server"
	"sl-monitor/internal/server/response"
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

	if err != nil || !users.CheckPasswordHash(creds.Pwd, idPass.Pwd) {
		response.Unauthorized(w, r)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)
	Sessions[sessionToken] = session{
		Email:  creds.Email,
		UserId: idPass.Id,
		Expiry: expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
}

func (ah *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			response.Unauthorized(w, r)
			return
		}
		response.BadRequest(w, r, err)
		return
	}
	sessionToken := c.Value

	delete(Sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}
