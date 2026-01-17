package repository

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/pkg/database"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type BookingRepo struct {
	db  database.PgxIface
	log *zap.Logger
}

type BookingRepoInterface interface {
	BookingSeat(ctx context.Context, tx pgx.Tx, bookingSeat *entity.BookingSeat) error
	UpdateSeatAvailability(ctx context.Context, tx pgx.Tx, seatID int) error
	GetIDByDateTime(ctx context.Context, seatID int, dateStr, timeStr string) (int, error)
	Payment(ctx context.Context, tx pgx.Tx, payment *entity.Payment) (*entity.Payment, error)
	BookingHistory(ctx context.Context, page, limit, userID int) ([]*entity.BookingHistory, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewBookingRepo(db database.PgxIface, log *zap.Logger) BookingRepoInterface {
	return &BookingRepo{
		db:  db,
		log: log,
	}
}

func (b *BookingRepo) BookingSeat(ctx context.Context, tx pgx.Tx, bookingSeat *entity.BookingSeat) error {
	query := `INSERT INTO booking_seat (user_id, showtime_id, seat_id, payment_method_id, status, created_at)
SELECT $1, $2, s.id, $4,'PENDING', $5
FROM seats s
WHERE s.id = $3 
AND s.is_available = TRUE 
RETURNING id;`

	now := time.Now()
	err := tx.QueryRow(ctx, query, bookingSeat.UserID, bookingSeat.ShowtimeId, bookingSeat.SeatId, bookingSeat.PaymentMethod, now).Scan(&bookingSeat.ID)
	if err != nil {
		b.log.Error("failed to create booking seat on database", zap.Error(err))
		return err
	}
	bookingSeat.CreatedAt = now
	return nil

}

func (b *BookingRepo) UpdateSeatAvailability(ctx context.Context, tx pgx.Tx, seatID int) error {
	query := `UPDATE seats SET is_available = false, updated_at = $1 WHERE id = $2`
	now := time.Now()
	_, err := tx.Exec(ctx, query, now, seatID)
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

func (b *BookingRepo) Payment(ctx context.Context, tx pgx.Tx, payment *entity.Payment) (*entity.Payment, error) {
	query := `UPDATE booking_seat
SET status='PAID', payment_details=$1
WHERE id=$2 AND user_id=$3 AND status='PENDING' AND payment_method_id=$4`
	// 1. Ubah struct details ke JSON String (Wajib untuk kolom JSONB/Text)
	detailsJSON, err := json.Marshal(payment.PaymentDetails)
	if err != nil {
		return nil, err
	}

	cmdTag, err := tx.Exec(ctx, query, detailsJSON, payment.BookingId, payment.UserID, payment.PaymentMethodID)
	if err != nil {
		b.log.Info("DEBUG PAYMENT",
			zap.Int("Input_BookingID", payment.BookingId),
			zap.Int("Input_UserID", payment.UserID),
			zap.Int("Input_MethodID", payment.PaymentMethodID),
		)
		return nil, err
	}
	if cmdTag.RowsAffected() == 0 {
		return nil, fmt.Errorf("payment failed: booking not found or status invalid")
	}

	return payment, nil
}

func (b *BookingRepo) BookingHistory(ctx context.Context, page, limit, userID int) ([]*entity.BookingHistory, error) {
	offset := (page - 1) * limit
	query := `SELECT m.title, m.duration, TO_CHAR(sh.start_time, 'YYYY-MM-DD HH24:MI') AS show_time, c.location
FROM booking_seat bs
JOIN showtimes sh ON bs.showtime_id = sh.id
JOIN movies m ON sh.movie_id = m.id
JOIN studios st ON sh.studio_id = st.id
JOIN cinemas c ON st.cinema_id = c.id
WHERE bs.user_id=$1
ORDER BY bs.created_at DESC
LIMIT $2 OFFSET $3;`

	rows, err := b.db.Query(ctx, query, userID, limit, offset)
	if err != nil {
		b.log.Error("failed to get booking history on database", zap.Error(err))
		return nil, err
	}
	var histories []*entity.BookingHistory

	for rows.Next() {
		var t entity.BookingHistory
		err := rows.Scan(&t.MovieTitle, &t.Duration, &t.ShowTime, &t.Location)
		if err != nil {
			b.log.Error("failed to scan rows", zap.Error(err))
		}
		histories = append(histories, &t)
	}
	return histories, nil
}

// method untuk memulai transaksi
func (b *BookingRepo) Begin(ctx context.Context) (pgx.Tx, error) {
	tx, err := b.db.Begin(ctx)
	if err != nil {
		b.log.Error("failed to create transaction", zap.Error(err))
		return nil, err
	}
	return tx, nil

}
