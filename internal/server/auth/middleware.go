package auth

import (
	"net/http"
	"sl-monitor/internal/server/response"
)

func MustBeLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		// We then get the name of the user from our session map, where we set the session token
		userSession, exists := sessions[sessionToken]
		if !exists {
			// If the session token is not present in session map, return an unauthorized error
			response.Unauthorized(w, r)
			return
		}
		if userSession.IsExpired() {
			delete(sessions, sessionToken)
			response.Unauthorized(w, r)
			return
		}
		next(w, r)
	}

}
