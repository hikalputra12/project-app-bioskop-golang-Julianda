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
