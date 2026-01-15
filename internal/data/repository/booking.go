package repository

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/pkg/database"
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
)

type BookingRepo struct {
	db  database.PgxIface
	log *zap.Logger
}

type BookingRepoInterface interface {
	BookingSeat(ctx context.Context, bookingSeat *entity.BookingSeat) error
	UpdateSeatAvailability(ctx context.Context, seat *entity.Seat) error
	GetIDByDateTime(ctx context.Context, seatID int, dateStr, timeStr string) (int, error)
}

func NewBookingRepo(db database.PgxIface, log *zap.Logger) BookingRepoInterface {
	return &BookingRepo{
		db:  db,
		log: log,
	}
}

func (b *BookingRepo) BookingSeat(ctx context.Context, bookingSeat *entity.BookingSeat) error {
	query := `INSERT INTO booking_seat (user_id, showtime_id, seat_id, payment_method_id, created_at)
SELECT $1, $2, s.id, $4, $5
FROM seats s
WHERE s.id = $3 
AND s.is_available = TRUE 
RETURNING id;`

	now := time.Now()
	err := b.db.QueryRow(ctx, query, bookingSeat.UserID, bookingSeat.ShowtimeId, bookingSeat.SeatId, bookingSeat.PaymentMethod, now).Scan(&bookingSeat.ID)
	if err != nil {
		b.log.Error("failed to create booking seat on database", zap.Error(err))
		return err
	}
	bookingSeat.CreatedAt = now
	return nil

}

func (b *BookingRepo) UpdateSeatAvailability(ctx context.Context, seat *entity.Seat) error {
	query := `UPDATE seats SET is_available = false, updated_at = $1 WHERE id = $2`
	now := time.Now()
	_, err := b.db.Exec(ctx, query, now, seat.ID)
	if err != nil {
		b.log.Error("failed to update seat availability on database", zap.Error(err))
		return err
	}
	return nil
}

func (b *BookingRepo) GetIDByDateTime(ctx context.Context, seatID int, dateStr, timeStr string) (int, error) {
	var showtimeID int

	query := `
        SELECT sh.id
        FROM showtimes sh
        JOIN seats st ON sh.studio_id = st.studio_id
        WHERE st.id = $1
      AND TO_CHAR(sh.start_time, 'YYYY-MM-DD') = $2
AND TO_CHAR(sh.start_time, 'HH24:MI') = $3;
  
    `

	err := b.db.QueryRow(ctx, query, seatID, dateStr, timeStr).Scan(&showtimeID)
	if err != nil {
		return 0, errors.New("jadwal film tidak ditemukan")
	}

	return showtimeID, nil
}
