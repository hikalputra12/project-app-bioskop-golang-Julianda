package adaptor

import (
	"app-bioskop/internal/dto"
	"app-bioskop/internal/usecase"
	"app-bioskop/pkg/utils"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type BookingAdaptor struct {
	bookingUsecase usecase.BookingUsecaseInterface
	sessionUsecase usecase.SessionUsecaseInterface
	log            *zap.Logger
}

func NewBookingAdaptor(bookingUsecase usecase.BookingUsecaseInterface, log *zap.Logger) *BookingAdaptor {
	return &BookingAdaptor{
		bookingUsecase: bookingUsecase,
		log:            log,
	}
}

// function booking seat adaptor
func (b *BookingAdaptor) BookingSeat(w http.ResponseWriter, r *http.Request) {
	var req dto.BookingRequest
	ctx := r.Context()

	//mengubah json body ke struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	userID := r.Context().Value("user_id").(int)
	err := b.bookingUsecase.BookingSeat(ctx, req, userID)
	if err != nil {
		b.log.Error("failed booking seat on service",
			zap.Error(err),
		)
		utils.ResponseError(w, http.StatusInternalServerError, "input tidak sesuai format yang di tentukan", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "booking seat succesfully",
	})
}

// // funtion to get booking history adaptor
// func (b *BookingAdaptor) BookingHistory(w http.ResponseWriter, r *http.Request) {
// 	page, err := strconv.Atoi(r.URL.Query().Get("page"))
// 	if err != nil {
// 		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid page", nil)
// 		return
// 	}

// 	// config limit pagination
// 	limit := 5
// 	ctx := r.Context()

// 	userID := r.Context().Value("user_id").(int)

// 	history, err := b.bookingUsecase.BookingHistory(ctx, page, limit, userID)
// 	if err != nil {
// 		b.log.Error("failed get booking history on service",
// 			zap.Error(err),
// 		)
// 		utils.ResponseError(w, http.StatusInternalServerError, "failed to get booking history", nil)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"status":  true,
// 		"message": "get booking history succesfully",
// 		"data":    history,
// 	})
// }

// // function payment adaptor
// func (b *BookingAdaptor) Payment(w http.ResponseWriter, r *http.Request) {
// 	var req dto.PaymentRequest
// 	ctx := r.Context()

// 	//mengubah json body ke struct
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid JSON format", nil)
// 		return
// 	}
// 	userID := r.Context().Value("user_id").(int)
// 	cmdTag, err := b.bookingUsecase.Payment(ctx, req, userID)
// 	if err != nil {
// 		b.log.Error("failed payment on service",
// 			zap.Error(err),
// 		)
// 		// Status 409 Conflict biasanya cocok untuk resource yang sudah diambil
// 		utils.ResponseError(w, http.StatusConflict, err.Error(), nil)
// 		return
// 	}
// 	if cmdTag == nil {
// 		utils.ResponseError(w, http.StatusBadRequest, "payment failed: booking not found or status invalid", nil)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"status":  true,
// 		"message": "payment succesfully",
// 	})
// }
