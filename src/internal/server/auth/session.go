package auth

import "time"

var Sessions = map[string]session{}

type session struct {
	Email  string
	Expiry time.Time
	UserId int
}

func (s session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
