package wire

import (
	"app-bioskop/internal/adaptor"
	"app-bioskop/internal/data/repository"
	"app-bioskop/internal/middleware"
	"app-bioskop/internal/usecase"
	"app-bioskop/pkg/utils"
	"sync"

	mid "github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func Wiring(repo *repository.Repository, log *zap.Logger, config utils.SMTPConfig) *chi.Mux {
	emailJobs := make(chan utils.EmailJob, 10) // BUFFER
	stop := make(chan struct{})
	metrics := &utils.Metrics{}
	wg := &sync.WaitGroup{}
	utils.StartEmailWorkers(3, emailJobs, stop, metrics, wg, config)
	useCase := usecase.AllUseCase(repo, log, emailJobs)
	adaptor := adaptor.AllAdaptor(useCase, log)

	r := chi.NewRouter()
	mw := middleware.MiddlewareCustome(useCase, log)
	r.Use(mw.Logging(log))
	r.Use(mid.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", adaptor.AuthAdaptor.Login)
		r.Post("/logout", adaptor.AuthAdaptor.Logout)
		r.Post("/register", adaptor.RegisterWire.Register)
		r.Post("/verify-otp", adaptor.VerifyAdaptor.SomeVerifyFunction)
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
