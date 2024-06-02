package ctxx

import (
	"context"
)

type ctxKey string

const (
	sessionID ctxKey = "sessions"
)

func SetSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, sessionID, sessionID)
}
func GetSessions(ctx context.Context) string {
	sessions, ok := ctx.Value(sessionID).(string)
	if !ok {
		return ""
	}
	return sessions
}
