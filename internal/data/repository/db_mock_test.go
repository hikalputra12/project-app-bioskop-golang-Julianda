package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// sqlmock untuk mock sql driver
// mock database gorm
func MockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	//pengecekan agar tidak terjadi error
	assert.NoError(t, err)

	mockDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db, DriverName: "postgres"}), &gorm.Config{})
	assert.NoError(t, err)
	return mockDB, mock

}
