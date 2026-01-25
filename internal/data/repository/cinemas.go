package repository

import (
	"app-bioskop/internal/data/entity"
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CinemaRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

type CinemasRepoInterface interface {
	GetAllCinemas(ctx context.Context, page, limit int) ([]*entity.Cinema, int64, error)
	GetCinemaByID(ctx context.Context, id int) (*entity.Cinema, error)
	GetSeatCinema(ctx context.Context, id int, date, time string) ([]*entity.Seat, error)
}

func NewCinemaRepo(db *gorm.DB, log *zap.Logger) CinemasRepoInterface {
	return &CinemaRepo{
		db:  db,
		log: log,
	}
}

// // get all cinemas
func (c *CinemaRepo) GetAllCinemas(ctx context.Context, page, limit int) ([]*entity.Cinema, int64, error) {
	var cinemas []*entity.Cinema
	// 	//menghitung offset
	offset := (page - 1) * limit
	// 	// get total data for pagination
	var total int64
	err := c.db.Model(&entity.Cinema{}).Count(&total)
	if err.Error != nil {
		return nil, 0, err.Error
	}

	result := c.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&cinemas)
	if result.Error != nil {
		panic("failed to connect to database to get all cinemas")
	}

	return cinemas, total, nil
}

// get cinemas by id
func (c *CinemaRepo) GetCinemaByID(ctx context.Context, id int) (*entity.Cinema, error) {
	var cinema entity.Cinema
	//menggunakan FInd akan menghasilkan sukses jika id tidak ada di bandingkan first akan langsung muncul error
	result := c.db.WithContext(ctx).First(&cinema, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cinema, nil

}

// // get seat cinema by date and time
func (c *CinemaRepo) GetSeatCinema(ctx context.Context, id int, date, timeStr string) ([]*entity.Seat, error) {
	//using GORM
	var seatCinema []*entity.Seat
	loc, _ := time.LoadLocation("Local")
	layout := "2006-01-02 15:04"

	startTime, err := time.ParseInLocation(layout, date+" "+timeStr, loc)
	if err != nil {
		return nil, err
	}
	result := c.db.Debug().WithContext(ctx).Model(&seatCinema).Select("seats.*").Joins("left join showtimes on showtimes.cinema_id=seats.cinema_id").
		Where("seats.cinema_id=? AND showtimes.start_time=?", id, startTime).Find(&seatCinema)
	if result.Error != nil {
		panic("failed connect to database to get cinema by id and time and date")
	}

	return seatCinema, nil
}
