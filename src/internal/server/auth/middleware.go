package auth

import (
	"net/http"
	"sl-monitor/internal/cache"
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

		userId := cache.Instance.FetchValue(sessionToken)
		if userId == "" {
			response.Unauthorized(w, r)
			return
		}
		next(w, r)
	}

}
