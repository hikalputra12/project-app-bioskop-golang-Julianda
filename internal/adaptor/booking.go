package adaptor

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/dto"
	"app-bioskop/internal/usecase"
	"app-bioskop/pkg/utils"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type BookingAdaptor struct {
	bookingUsecase usecase.BookingUsecaseInterface
	log            *zap.Logger
}

func NewBookingAdaptor(bookingUsecase usecase.BookingUsecaseInterface, log *zap.Logger) *BookingAdaptor {
	return &BookingAdaptor{
		bookingUsecase: bookingUsecase,
		log:            log,
	}
}

func (b *BookingAdaptor) BookingSeat(w http.ResponseWriter, r *http.Request) {
	var req dto.BookingRequest
	ctx := r.Context()

	//mengubah json body ke struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	for _, seatID := range req.SeatIds {
		showtimeID, err := b.bookingUsecase.GetIDByDateTime(ctx, seatID, req.Date, req.Time)
		if err != nil {
			b.log.Error("failed get showtime id by date time on service",
				zap.Error(err),
			)
			utils.ResponseError(w, http.StatusBadRequest, "input tidak sesuai format yang di tentukan", nil)
			return
		}
		newBookingSeat := entity.BookingSeat{
			UserID:        req.UserID,
			ShowtimeId:    showtimeID,
			SeatId:        seatID,
			PaymentMethod: req.PaymentMethod,
		}
		err = b.bookingUsecase.BookingSeat(ctx, &newBookingSeat)
		if err != nil {
			b.log.Error("failed booking seat on service",
				zap.Error(err),
			)
			utils.ResponseError(w, http.StatusBadRequest, "input tidak sesuai format yang di tentukan", nil)
			return
		}
		seat := entity.Seat{
			Entity: entity.Entity{
				ID: seatID,
			},
		}
		err = b.bookingUsecase.UpdateSeatAvailability(ctx, &seat)
		if err != nil {
			b.log.Error("failed update seat availability on service",
				zap.Error(err),
			)
			utils.ResponseError(w, http.StatusBadRequest, "input tidak sesuai format yang di tentukan", nil)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "booking seat succesfully",
	})
}
