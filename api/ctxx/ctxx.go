package ctxx

import (
	"context"

	"github.com/gorilla/sessions"
)

type ctxKey string

const (
	sessionID ctxKey = "sessions"
	csrfToken ctxKey = "csrfToken"
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

func SetCSRFToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, csrfToken, token)
}
func GetCSRFToken(ctx context.Context) string {
	token, ok := ctx.Value(csrfToken).(string)
	if !ok {
		return ""
	}
	return token
}
