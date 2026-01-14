package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Logging adalah middleware untuk mencatat setiap request HTTP
func (m *Middleware) Logging(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Catat log setelah request selesai
			duration := time.Since(start)
			// Log Request Masuk
			m.log.Info("Incoming Request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
			)

			// Lanjut ke handler berikutnya
			next.ServeHTTP(w, r)

			// Log Durasi Selesai
			m.log.Info("Request Completed",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Duration("latency", duration),
			)
		})
	}
}
