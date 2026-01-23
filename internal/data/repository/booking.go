package repository

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/dto"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BookingRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

type BookingRepoInterface interface {
	BookingAndUpdateSeat(ctx context.Context, bookingSeat *entity.BookingSeat, seatID int) error
	GetShowtimeID(ctx context.Context, seatID int, dateStr, timeStr string) (int, error)
	Payment(ctx context.Context, payment *entity.BookingSeat) error
	BookingHistory(ctx context.Context, page, limit, userID int) ([]*dto.BookingHistoryResponse, error)
}

func NewBookingRepo(db *gorm.DB, log *zap.Logger) BookingRepoInterface {
	return &BookingRepo{
		db:  db,
		log: log,
	}
}

// function to book seat
func (b *BookingRepo) BookingAndUpdateSeat(ctx context.Context, bookingSeat *entity.BookingSeat, seatID int) error {
	var seat *entity.Seat
	b.db.Transaction(func(tx *gorm.DB) error {
		result := tx.WithContext(ctx).Model(&seat).Where("id=? AND is_available=?", bookingSeat.SeatId, true).Update("is_available", false)
		if result.Error != nil {
			b.log.Error("error database saat update kursi", zap.Error(result.Error))
			return result.Error
		}
		if result.RowsAffected == 0 {
			b.log.Warn("gagal karena kursi sudha di ambil")
			return fmt.Errorf("kursi sudah tidak tersedia")
		}

		tx.Transaction(func(tx *gorm.DB) error {
			if err := tx.WithContext(ctx).Create(&bookingSeat).Error; err != nil {
				return err
			}
			return nil
		})
		return nil

	})

	return nil

}

func (b *BookingRepo) GetShowtimeID(ctx context.Context, seatID int, dateStr, timeStr string) (int, error) {

	loc, _ := time.LoadLocation("Local") // Sesuaikan timezone
	layout := "2006-01-02 15:04"
	targetTime, err := time.ParseInLocation(layout, dateStr+" "+timeStr, loc)
	if err != nil {
		return 0, fmt.Errorf("invalid date time format")
	}

	var showtimeID int

	// Cari Showtime ID
	err = b.db.WithContext(ctx).
		Table("showtimes").
		Select("showtimes.id").
		Joins("JOIN seats ON seats.cinema_id = showtimes.cinema_id").
		Where("seats.id = ? AND showtimes.start_time = ?", seatID, targetTime).
		Scan(&showtimeID).Error

	if err != nil {
		return 0, err
	}

	if showtimeID == 0 {
		return 0, fmt.Errorf("showtime not found")
	}

	return showtimeID, nil
}

func (b *BookingRepo) Payment(ctx context.Context, payment *entity.BookingSeat) error {
	b.db.Transaction(func(tx *gorm.DB) error {
		result := tx.WithContext(ctx).Model(&payment).Where("id=? AND payment_method_id=?", payment.ID, payment.PaymentMethodID).Updates(map[string]interface{}{"payment_details": payment.PaymentDetails, "status": "PAID"})
		if result.Error != nil {
			b.log.Warn("gagal melakukan pembayaran")
			return result.Error
		}
		return nil
	})
	return nil
}

// // function booking history repository
func (b *BookingRepo) BookingHistory(ctx context.Context, page, limit, userID int) ([]*dto.BookingHistoryResponse, error) {
	offset := (page - 1) * limit
	var bookingHistory []*dto.BookingHistoryResponse
	var booking []*entity.BookingSeat

	result := b.db.WithContext(ctx).
		Joins("JOIN showtimes ON showtimes.id = booking_seats.showtime_id").
		Joins("JOIN movies ON movies.id = showtimes.movie_id").
		Joins("JOIN cinemas ON cinemas.id = showtimes.cinema_id").Select("movies.title AS movie_title,movies.duration,showtimes.start_time AS show_time,cinemas.name AS cinema_name,cinemas.location").
		Where("booking_seats.user_id = ? AND booking_seats.status = ?", userID, "PAID").
		Limit(limit).
		Offset(offset).Find(&booking).
		Scan(&bookingHistory)
	if result.Error != nil {
		b.log.Warn("gagal mengambil booking history")
		return nil, result.Error
	}

	return bookingHistory, nil
}
