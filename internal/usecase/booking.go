package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"app-bioskop/internal/dto"
	"context"

	"go.uber.org/zap"
)

type BookingUsecase struct {
	bookingUsecase repository.BookingRepoInterface
	log            *zap.Logger
}

type BookingUsecaseInterface interface {
	BookingSeat(ctx context.Context, req dto.BookingRequest, userID int) error
	Payment(ctx context.Context, req dto.PaymentRequest, userID int) error
	BookingHistory(ctx context.Context, page, limit, userID int) ([]*dto.BookingHistoryResponse, error)
}

func NewBookingUsecase(bookingUsecase repository.BookingRepoInterface, log *zap.Logger) BookingUsecaseInterface {
	return &BookingUsecase{
		bookingUsecase: bookingUsecase,
		log:            log,
	}
}

// logic to booking seat
func (b *BookingUsecase) BookingSeat(ctx context.Context, req dto.BookingRequest, userID int) error {

	for _, seatID := range req.SeatIds {
		showtimeID, err := b.bookingUsecase.GetShowtimeID(ctx, seatID, req.Date, req.Time)
		if err != nil {
			b.log.Error("failed get showtime id by date time on repo",
				zap.Error(err),
			)
		}
		newBookingSeat := entity.BookingSeat{
			UserID:          userID,
			ShowTimeId:      showtimeID,
			SeatId:          seatID,
			PaymentMethodID: req.PaymentMethod,
			Status:          "PENDING",
		}

		//booking seat
		err = b.bookingUsecase.BookingAndUpdateSeat(ctx, &newBookingSeat, seatID)
		if err != nil {
			b.log.Error("failed booking seat on service",
				zap.Error(err),
			)
			return err
		}
	}

	return nil
}

// // logic to get booking history
func (b *BookingUsecase) BookingHistory(ctx context.Context, page, limit, userID int) ([]*dto.BookingHistoryResponse, error) {
	history, err := b.bookingUsecase.BookingHistory(ctx, page, limit, userID)
	if err != nil {
		b.log.Error("failed to get booking history on repository", zap.Error(err))
		return nil, err
	}
	return history, nil
}

// // logic to payment
func (b *BookingUsecase) Payment(ctx context.Context, req dto.PaymentRequest, userID int) error {

	payment := &entity.BookingSeat{
		UserID: userID,
		Entity: entity.Entity{
			ID: req.BookingId,
		},
		PaymentMethodID: req.PaymentMethod,
		PaymentDetails: entity.PaymentDetails{
			CardNumber: req.PaymentDetails.CardNumber,
			CVV:        req.PaymentDetails.CVV,
			ExpiryDate: req.PaymentDetails.ExpiryDate,
		},
	}

	//process payment
	err := b.bookingUsecase.Payment(ctx, payment)
	if err != nil {
		b.log.Error("failed to process payment on repository", zap.Error(err))
		return err
	}

	return nil
}
