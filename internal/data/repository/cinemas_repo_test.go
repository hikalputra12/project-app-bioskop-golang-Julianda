package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCinema struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	repo   Repository
	gormDB *gorm.DB
}

func (s *suiteCinema) SetupTest() {
	s.gormDB, s.mock = MockDB(s.T())

	s.repo.CinemaRepo = NewCinemaRepo(s.gormDB, nil)
}

func (s *suiteCinema) TestGetAllCinemas() {

	s.Run("success", func() {
		// Data Dummy
		rows := sqlmock.NewRows([]string{
			"id", "name", "location", "deleted_at",
		}).AddRow(1, "XX1", "Mall Panakukang", nil)

		countRow := sqlmock.NewRows([]string{"count"}).AddRow(1)

		s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "cinemas"`)).
			WillReturnRows(countRow)

		//menggunakan regexp.QuoteMeta untuk membaca *
		s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cinemas"`)).
			WillReturnRows(rows)

		cinemas, total, err := s.repo.CinemaRepo.GetAllCinemas(context.Background(), 1, 10)

		//lakukan validasi
		s.NoError(err)
		s.NotNil(cinemas)
		s.Equal(int64(1), total)
		s.Equal("XX1", cinemas[0].Name)
		s.Equal("Mall Panakukang", cinemas[0].Location)

		// Cek apakah semua urutan mock terpenuhi
		s.NoError(s.mock.ExpectationsWereMet())
	})

	s.Run("not_found", func() {
		s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "cinemas"`)).
			WillReturnError(gorm.ErrRecordNotFound)

		cinemas, total, err := s.repo.CinemaRepo.GetAllCinemas(context.Background(), 1, 10)

		s.Error(err)
		s.Nil(cinemas)
		s.Equal(int64(0), total)

		s.NoError(s.mock.ExpectationsWereMet())
	})
}

func (s *suiteCinema) TestGetCinemasByID() {

	s.Run("success", func() {
		id := 1
		// Data Dummy
		rows := sqlmock.NewRows([]string{
			"id", "name", "location", "deleted_at",
		}).AddRow(id, "Cinepolis", "Mall Panakukang", nil)

		//menggunakan regexp.QuoteMeta untuk membaca *
		s.mock.ExpectQuery(`FROM "cinemas"`).WithArgs(id, sqlmock.AnyArg()).
			WillReturnRows(rows)

		cinema, err := s.repo.CinemaRepo.GetCinemaByID(context.Background(), id)

		//lakukan validasi
		s.NoError(err)
		s.NotNil(cinema)
		s.Equal(id, cinema.ID)
		s.Equal("Cinepolis", cinema.Name)
		s.Equal("Mall Panakukang", cinema.Location)

		// Cek apakah semua urutan mock terpenuhi
		s.NoError(s.mock.ExpectationsWereMet())
	})

	s.Run("not_found", func() {
		id := 2
		s.mock.ExpectQuery(`FROM "cinemas"`).WithArgs(id, sqlmock.AnyArg()).
			WillReturnError(gorm.ErrRecordNotFound)

		cinema, err := s.repo.CinemaRepo.GetCinemaByID(context.Background(), id)

		s.Error(err)
		s.Nil(cinema)

		s.NoError(s.mock.ExpectationsWereMet())
	})
}

func (s *suiteCinema) TestGetSeatCinema() {

	s.Run("success", func() {
		id := 1
		dateStr := "2006-01-02"
		timeStr := "15:40"

		// Data Dummy
		rows := sqlmock.NewRows([]string{"id",
			"seat_number", "is_available", "deleted_at",
		}).AddRow(id, "A1", true, nil)

		//menggunakan regexp.QuoteMeta untuk membaca *
		s.mock.ExpectQuery(`FROM "seats"`).WithArgs(id, sqlmock.AnyArg()).
			WillReturnRows(rows)

		seats, err := s.repo.CinemaRepo.GetSeatCinema(context.Background(), id, dateStr, timeStr)

		//lakukan validasi
		s.NoError(err)
		s.NotNil(seats)
		s.Require().NotEmpty(seats, "Data tidak boleh kosong")
		s.Equal(id, seats[0].ID, "ID yang dikembalikan harus sama dengan yang diminta")
		s.NoError(s.mock.ExpectationsWereMet())
	})

	s.Run("not_found", func() {
		id := 3
		dateStr := "2009-01-02"
		timeStr := "16:40"

		s.mock.ExpectQuery(`FROM "seats"`).WithArgs(id, sqlmock.AnyArg()).
			WillReturnError(gorm.ErrRecordNotFound)

		seats, err := s.repo.CinemaRepo.GetSeatCinema(context.Background(), id, dateStr, timeStr)

		s.Error(err)
		s.Nil(seats)

		s.NoError(s.mock.ExpectationsWereMet())
	})
}

// fungsi untuk test semua fungsi testing yang di kumpulkan di suite
func TestCinemaRepo(t *testing.T) {
	suite.Run(t, new(suiteCinema))
}
