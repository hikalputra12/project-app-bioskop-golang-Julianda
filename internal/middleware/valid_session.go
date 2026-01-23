package middleware

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/pkg/utils"
	"context"
	"net/http"
	"time"
)

func (m *Middleware) ValidExtend() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//ambil cookie
			ctx := r.Context()
			session, err := r.Cookie("session")

			if err != nil {
				utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized: sesi tidak ditemukan", nil)
				return
			}
			//ambil value dari cookie
			getSessionID := session.Value
			//set user id di context
			userID, err := m.Usecase.SessionUsecase.GetUserIDBySession(ctx, getSessionID)

			ctxUserId := context.WithValue(r.Context(), "user_id", userID)

			sessionID := &entity.Session{
				ID: getSessionID,
			}

			//cek apakah valid atau tidak dan cek apakah ada cookie atau tidak dari error
			valid, err := m.Usecase.SessionUsecase.IsValid(ctx, sessionID)
			if err != nil {
				return
			}
			if !valid {
				utils.ResponseError(w, http.StatusUnauthorized, "Sesi sudah kadaluarsa atau tidak valid", nil)
				return
			}
			//set uuid baru
			newSession := utils.NewUUID()
			//exted cookie
			NewsessionID := &entity.Session{
				NewID:      newSession,
				ID:         getSessionID,
				ExpiredAt:  time.Now().Add(24 * time.Hour),
				LastActive: time.Now(),
			}
			err = m.Usecase.SessionUsecase.ExtendSession(ctx, NewsessionID)
			if err != nil {
				return
			}

			//perbaharui cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "session",
				Value:    newSession,
				Path:     "/",
				MaxAge:   24 * 60 * 60,
				HttpOnly: true,
				Secure:   false,
			})

			next.ServeHTTP(w, r.WithContext(ctxUserId))
		})
	}
}
