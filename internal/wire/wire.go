package wire

import (
	"app-bioskop/internal/adaptor"
	"app-bioskop/internal/data/repository"
	"app-bioskop/internal/middleware"
	"app-bioskop/internal/usecase"

	mid "github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Wiring(repo *repository.Repository, log *zap.Logger) *chi.Mux {

	useCase := usecase.AllUseCase(repo, log)
	adaptor := adaptor.AllAdaptor(useCase, log)

	r := chi.NewRouter()
	mw := middleware.MiddlewareCustome(useCase, log)
	r.Use(mw.Logging(log))
	r.Use(mid.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", adaptor.AuthAdaptor.Login)
		r.Post("/logout", adaptor.AuthAdaptor.Logout)
		r.Post("/register", adaptor.RegisterWire.Register)
		r.Get("/cinemas", adaptor.CinemaAdaptor.GetAllCinemas)
		r.Get("/cinemas/{cinemaId}", adaptor.CinemaAdaptor.GetcinemasById)
		r.Get("/cinemas/{cinemaId}/seats", adaptor.CinemaAdaptor.GetSeatCinema)
		r.Get("/payment-methods", adaptor.PaymentMethodAdaptor.GetAllPaymentMethods)
		r.With(mw.ValidExtend()).Post("/booking", adaptor.BookingAdaptor.BookingSeat)
		r.With(mw.ValidExtend()).Post("/payment", adaptor.BookingAdaptor.Payment)
		r.With(mw.ValidExtend()).Get("/user/bookings", adaptor.BookingAdaptor.BookingHistory)
	})

	return r
}
