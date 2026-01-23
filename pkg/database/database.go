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

// type PgxIface interface {
// 	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
// 	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
// 	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
// 	Begin(ctx context.Context) (pgx.Tx, error)
// }

// func InitDB(config utils.DatabaseCofig) (*pgxpool.Pool, error) {
// 	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s",
// 		config.Username, config.Password, config.Name, config.Host)
// cfg, err := pgxpool.ParseConfig(connStr)
// if err != nil {
// 	return nil, fmt.Errorf("parse config: %w", err)
// }

// cfg.MaxConns = config.MaxConn
// cfg.MinConns = 5
// cfg.MaxConnLifetime = 30 * time.Minute
// cfg.MaxConnIdleTime = 5 * time.Minute
// cfg.HealthCheckPeriod = 1 * time.Minute

// cfg.ConnConfig.ConnectTimeout = 5 * time.Second

// pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
// if err != nil {
// 	return nil, fmt.Errorf("create pool: %w", err)
// }

// pingCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// defer cancel()
// if err := pool.Ping(pingCtx); err != nil {
// 	pool.Close()
// 	return nil, fmt.Errorf("ping db: %w", err)
// }

// return pool, nil
// }

//belajar untuk menghubungkan gorm.DB dengan postgress

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
