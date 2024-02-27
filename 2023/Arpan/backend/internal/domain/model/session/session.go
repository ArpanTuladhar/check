package session

import (
	"context"
	"errors"
)

type CtxKeySession struct{}
type Session struct {
	UserId int32
}

func WithSession(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, &CtxKeySession{}, session)
}

func ExtractSession(ctx context.Context) (*Session, error) {
	s, ok := ctx.Value(&CtxKeySession{}).(*Session)
	if ok {
		return s, nil
	}
	return nil, errors.New("session not found")
}

func StoreSession(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, &CtxKeySession{}, session)
}
