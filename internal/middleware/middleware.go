package middleware

import (
	"app-bioskop/internal/usecase"

	"go.uber.org/zap"
)

type Middleware struct {
	Usecase usecase.Usecase
	log     *zap.Logger
}

func MiddlewareCustome(Usecase usecase.Usecase, log *zap.Logger) Middleware {
	return Middleware{
		Usecase: Usecase,
		log:     log,
	}
}
