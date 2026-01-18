package usecase

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/internal/data/repository"
	"app-bioskop/internal/dto"
	"app-bioskop/pkg/utils"
	"context"

	"go.uber.org/zap"
)

type CinemaUsecase struct {
	cinemaRepo repository.CinemasRepoInterface
	log        *zap.Logger
}

type CinemaUsecaseInterface interface {
	GetAllCinemas(ctx context.Context, page, limit int) ([]*entity.Cinema, *dto.Pagination, error)
	GetCinemaByID(ctx context.Context, id int) (*entity.Cinema, error)
	GetSeatCinema(ctx context.Context, id int, date, time string) ([]*entity.Seat, error)
}

// create new cinema usecase
func NewCinemaUsecase(cinemaRepo repository.CinemasRepoInterface, log *zap.Logger) CinemaUsecaseInterface {
	return &CinemaUsecase{
		cinemaRepo: cinemaRepo,
		log:        log,
	}
}

// get all cinemas
func (c *CinemaUsecase) GetAllCinemas(ctx context.Context, page, limit int) ([]*entity.Cinema, *dto.Pagination, error) {
	rows, total, err := c.cinemaRepo.GetAllCinemas(ctx, page, limit)
	if err != nil {
		c.log.Error("failed get all cinemas on repository ",
			zap.Error(err),
		)
		return nil, nil, err
	}
	pagination := dto.Pagination{
		CurrentPage: page,
		Limit:       limit,
		TotalPages:  utils.TotalPage(limit, int64(total)),
	}
	return rows, &pagination, nil
}

// get cinemas by id
func (c *CinemaUsecase) GetCinemaByID(ctx context.Context, id int) (*entity.Cinema, error) {
	row, err := c.cinemaRepo.GetCinemaByID(ctx, id)
	if err != nil {
		c.log.Error("failed get cinemas by id on repository ",
			zap.Error(err),
		)
		return nil, err
	}
	return row, nil
}

// get seat cinema by date and time
func (c *CinemaUsecase) GetSeatCinema(ctx context.Context, id int, date, time string) ([]*entity.Seat, error) {
	rows, err := c.cinemaRepo.GetSeatCinema(ctx, id, date, time)
	if err != nil {
		c.log.Error("failed get all seat cinema by id cinema on repository ",
			zap.Error(err),
		)
		return nil, err
	}

	return rows, nil
}
