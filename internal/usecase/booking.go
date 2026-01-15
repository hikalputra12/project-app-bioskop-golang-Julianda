package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"context"

	"go.uber.org/zap"
)

type BookingUsecase struct {
	bookingUsecase repository.BookingRepoInterface
	log            *zap.Logger
}

type BookingUsecaseInterface interface {
	BookingSeat(ctx context.Context, bookingSeat *entity.BookingSeat) error
	UpdateSeatAvailability(ctx context.Context, seat *entity.Seat) error
	GetIDByDateTime(ctx context.Context, seatID int, dateStr, timeStr string) (int, error)
}

func NewBookingUsecase(bookingUsecase repository.BookingRepoInterface, log *zap.Logger) BookingUsecaseInterface {
	return &BookingUsecase{
		bookingUsecase: bookingUsecase,
		log:            log,
	}
}

func (b *BookingUsecase) BookingSeat(ctx context.Context, bookingSeat *entity.BookingSeat) error {
	err := b.bookingUsecase.BookingSeat(ctx, bookingSeat)
	if err != nil {
		b.log.Error("failed to booking seat on repository", zap.Error(err))
		return err
	}
	return nil
}

func (b *BookingUsecase) UpdateSeatAvailability(ctx context.Context, seat *entity.Seat) error {
	err := b.bookingUsecase.UpdateSeatAvailability(ctx, seat)
	if err != nil {
		b.log.Error("failed to update seat availability on repository", zap.Error(err))
		return err
	}
	return nil
}

func (b *BookingUsecase) GetIDByDateTime(ctx context.Context, seatID int, dateStr, timeStr string) (int, error) {
	showtimeID, err := b.bookingUsecase.GetIDByDateTime(ctx, seatID, dateStr, timeStr)
	if err != nil {
		b.log.Error("failed to get showtime id by date time on repository", zap.Error(err))
		return 0, err
	}
	return showtimeID, nil
}
