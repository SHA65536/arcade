package session

import "github.com/google/uuid"

type Session struct {
	SessionId uuid.UUID
	Username  string
}

func NewSession() *Session {
	return &Session{
		SessionId: uuid.New(),
		Username:  "Guest",
	}
}
