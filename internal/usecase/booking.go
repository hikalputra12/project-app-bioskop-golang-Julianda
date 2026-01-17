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
	Payment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error)
	BookingHistory(ctx context.Context, page, limit, userID int) ([]*entity.BookingHistory, error)
}

func NewBookingUsecase(bookingUsecase repository.BookingRepoInterface, log *zap.Logger) BookingUsecaseInterface {
	return &BookingUsecase{
		bookingUsecase: bookingUsecase,
		log:            log,
	}
}

func (b *BookingUsecase) BookingSeat(ctx context.Context, req dto.BookingRequest, userID int) error {
	tx, err := b.bookingUsecase.Begin(ctx)
	if err != nil {
		return err
	}
	//untuk menghentikan transaksi jika gagal
	defer tx.Rollback(ctx)
	for _, seatID := range req.SeatIds {
		showtimeID, err := b.bookingUsecase.GetIDByDateTime(ctx, seatID, req.Date, req.Time)
		if err != nil {
			b.log.Error("failed get showtime id by date time on repo",
				zap.Error(err),
			)
		}
		newBookingSeat := entity.BookingSeat{
			UserID:        userID,
			ShowtimeId:    showtimeID,
			SeatId:        seatID,
			PaymentMethod: req.PaymentMethod,
		}
		//booking seat
		err = b.bookingUsecase.BookingSeat(ctx, tx, &newBookingSeat)
		if err != nil {
			b.log.Error("failed booking seat on service",
				zap.Error(err),
			)
		}
		//update seat to database
		err = b.bookingUsecase.UpdateSeatAvailability(ctx, tx, seatID)
		if err != nil {
			b.log.Error("failed update seat availability on service",
				zap.Error(err),
			)
		}
	}
	// Commit: meyimpan perubahan permanen jika semua loop sukses
	if err := tx.Commit(ctx); err != nil {
		b.log.Error("failed to commit transaction", zap.Error(err))
		return err
	}

	return nil
}

func (b *BookingUsecase) BookingHistory(ctx context.Context, page, limit, userID int) ([]*entity.BookingHistory, error) {
	history, err := b.bookingUsecase.BookingHistory(ctx, page, limit, userID)
	if err != nil {
		b.log.Error("failed to get booking history on repository", zap.Error(err))
		return nil, err
	}
	return history, nil
}

func (b *BookingUsecase) Payment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	pay, err := b.bookingUsecase.Payment(ctx, payment)
	if err != nil {
		b.log.Error("failed to process payment on repository", zap.Error(err))
		return nil, err
	}
	return pay, nil
}
