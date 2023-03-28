package auth

import (
	"github.com/google/uuid"
	"net/http"
	"sl-monitor/internal/config"
	"sl-monitor/internal/server"
	"sl-monitor/internal/server/response"
	"time"
)

type Handler struct {
	config *config.Config
}

func NewHandler(config *config.Config) *Handler {
	return &Handler{config}
}

var users = map[string]string{ // TODO get rid
	"user1": "password1",
	"user2": "password2",
}

func (ah *Handler) login(w http.ResponseWriter, r *http.Request) {
	type Credentials struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}

	var creds Credentials

	err := request.DecodeJSON(r, &creds)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		response.Unauthorized(w, r)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)
	Sessions[sessionToken] = session{
		Username: creds.Username,
		UserId:   123, // TODO
		Expiry:   expiresAt,
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
