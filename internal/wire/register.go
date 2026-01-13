package wire

import (
	"app-bioskop/internal/adaptor"
	"app-bioskop/internal/data/repository"
	"app-bioskop/internal/usecase"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Wiring(repo repository.Repository, log *zap.Logger) *chi.Mux {

	useCase := usecase.AllUseCase(repo, log)
	handler := adaptor.AllAdaptor(useCase, log)

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Post("/register", handler.RegisterWire.Register)

	})

	return r
}
