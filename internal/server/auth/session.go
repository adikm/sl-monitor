package auth

import "time"

var sessions = map[string]session{}

type session struct {
	Username string
	Expiry   time.Time
}

func (s session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
