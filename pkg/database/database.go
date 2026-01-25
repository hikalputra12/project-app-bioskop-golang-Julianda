package database

import (
	"app-bioskop/pkg/utils"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB(config utils.DatabaseCofig) (*gorm.DB, error) {

	//logger for gorm
	logger := logger.New(log.New(os.Stdout, "r/n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		})

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s TimeZone=Asia/Jakarta", config.Username, config.Password, config.Name, config.Host)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "public."},
		Logger:         logger,
	})
	if err != nil {
		panic("can't connect to database")
	}
	return db, nil
}
