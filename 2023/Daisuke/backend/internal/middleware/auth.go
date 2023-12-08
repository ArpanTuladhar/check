package middleware

import (
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/model/session"
	"net/http"
)

func (m middleware) WithAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//TODO: confirm session
			s := session.Session{UserId: 12345}
			ctx := session.StoreSession(r.Context(), &s)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
