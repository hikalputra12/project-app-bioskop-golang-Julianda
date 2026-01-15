package repository

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/pkg/database"
	"context"

	"go.uber.org/zap"
)

type CinemaRepo struct {
	db  database.PgxIface
	log *zap.Logger
}

type CinemasRepoInterface interface {
	GetAllCinemas(ctx context.Context, page, limit int) ([]*entity.Cinema, int, error)
	GetCinemaByID(ctx context.Context, id int) (*entity.Cinema, error)
	GetSeatCinema(ctx context.Context, id int, date, time string) ([]*entity.Seat, error)
}

func NewCinemaRepo(db database.PgxIface, log *zap.Logger) CinemasRepoInterface {
	return &CinemaRepo{
		db:  db,
		log: log,
	}
}

func (c *CinemaRepo) GetAllCinemas(ctx context.Context, page, limit int) ([]*entity.Cinema, int, error) {
	//menghitung offset
	offset := (page - 1) * limit
	// get total data for pagination
	var total int
	countQuery := `SELECT COUNT(*) FROM cinemas WHERE deleted_at IS NULL`
	err := c.db.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT name,location from cinemas ORDER BY id ASC LIMIT $1 OFFSET $2;`
	rows, err := c.db.Query(ctx, query, limit, offset)
	if err != nil {
		c.log.Error("Database Query Error: failed get all cinemas on database",
			zap.Error(err),
			zap.String("query", query),
		)
		return nil, 0, err
	}
	defer rows.Close()

	var cinemas []*entity.Cinema

	for rows.Next() {
		var t entity.Cinema
		err := rows.Scan(&t.Name, &t.Location)
		if err != nil {
			return nil, 0, err
		}
		cinemas = append(cinemas, &t)
	}
	return cinemas, total, nil
}

func (c *CinemaRepo) GetCinemaByID(ctx context.Context, id int) (*entity.Cinema, error) {
	var cinemas entity.Cinema
	query := `SELECT name,location from cinemas WHERE id=$1;`
	err := c.db.QueryRow(ctx, query, id).Scan(&cinemas.Name, &cinemas.Location)
	if err != nil {
		c.log.Error("failed to get cinema by id on database", zap.Error(err), zap.String("query", query))
		return nil, err
	}
	return &cinemas, nil
}

func (c *CinemaRepo) GetSeatCinema(ctx context.Context, id int, date, time string) ([]*entity.Seat, error) {
	query := `SELECT s.seat_number, s.is_available FROM seats s 
JOIN studios st ON st.id=s.studio_id
JOIN cinemas c ON c.id= st.cinema_id
JOIN showtimes sh ON sh.studio_id = st.id
WHERE c.id=$1
AND TO_CHAR(sh.start_time, 'YYYY-MM-DD') = $2 
AND TO_CHAR(sh.start_time, 'HH24:MI') = $3;

`
	rows, err := c.db.Query(ctx, query, id, date, time)
	if err != nil {
		c.log.Error("failed to get seat by id on database", zap.Error(err), zap.String("query", query))
		return nil, err
	}
	var seats []*entity.Seat
	for rows.Next() {
		var t entity.Seat
		err := rows.Scan(&t.SeatNumber, &t.IsAvaiable)
		if err != nil {
			return nil, err
		}
		seats = append(seats, &t)
	}

	return seats, nil
}
