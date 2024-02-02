package session

import (
	"context"
	"errors"
)

type ctxKeySession struct{}

type Session struct {
	UserId int32
}

func ExtractSession(ctx context.Context) (*Session, error) {
	s, ok := ctx.Value(&ctxKeySession{}).(*Session)
	if ok {
		return s, nil
	}
	return nil, errors.New("session not found")
}

func StoreSession(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, &ctxKeySession{}, session)
}
