package ctxx

import (
	"context"

	"github.com/gorilla/sessions"
)

type ctxKey string

const (
	sessionID ctxKey = "sessions"
)

func SetSessions(ctx context.Context, sessions *sessions.Session) context.Context {
	return context.WithValue(ctx, sessionID, sessions)
}

func GetSessions(ctx context.Context) *sessions.Session {
	sessions, ok := ctx.Value(sessionID).(*sessions.Session)
	if !ok {
		return nil
	}
	return sessions
}
