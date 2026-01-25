package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// buat struct suite agar teratur
type Suite struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	gormDB *gorm.DB
	repo   Repository
}

func (s *Suite) SetupSuite() {
	s.gormDB, s.mock = MockDB(s.T()) //masukkan testing context

	s.repo.AuthRepo = NewAuthRepo(s.gormDB, nil)
}

func (s *Suite) TestFindByEmail() {

	//buat data dummy
	s.Run("success", func() {
		email := "siapa@gmail.com"
		rows := sqlmock.NewRows([]string{
			"id", "email", "password_hash", "deleted_at",
		}).AddRow(1, email, "password", nil)

		//ekspetasi query yang akan di hasilkan
		//sqlmock.AnyArg() artunya argumen apapun itu
		s.mock.ExpectQuery(`FROM "users"`).WithArgs(email, sqlmock.AnyArg()).WillReturnRows(rows)

		//pemanggilan repo
		user, err := s.repo.AuthRepo.FindByEmail(context.Background(), email)
		s.NoError(err)
		s.NotNil(user)
		s.Equal(email, user.Email)
		s.Equal("password", user.Password)
		s.NoError(s.mock.ExpectationsWereMet())
		s.T().Log("Testing FindByEmail success case")
	})
	//ekspetasi email tidak di temukan
	s.Run("not_found", func() {
		email := "tidakada@gmail.com"
		s.mock.ExpectQuery(`FROM "users`).
			WithArgs(email, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		user, err := s.repo.AuthRepo.FindByEmail(context.Background(), email)
		s.Error(err)
		s.Nil(user)
	})
}

//fungsi untuk test semua fungsi testing yang di kumpulkan di suite

func TestAuthRepo(t *testing.T) {
	suite.Run(t, new(Suite))
}
