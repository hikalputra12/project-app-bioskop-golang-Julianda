package repository

import (
	"app-bioskop/pkg/database"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type Verify struct {
	db  database.PgxIface
	log *zap.Logger
}

type VerifyInterface interface {
	VerifyOTP(ctx context.Context, email string) error
}

func NewVerifyRepo(db database.PgxIface, log *zap.Logger) VerifyInterface {
	return &Verify{
		db:  db,
		log: log,
	}
}

// VerifyOTP updates the user's is_verified status to true based on the provided email.
func (b *Verify) VerifyOTP(ctx context.Context, email string) error {

	query := `UPDATE users
SET is_verified = true WHERE email = $1`
	cmdTag, err := b.db.Exec(ctx, query, email)
	if err != nil {
		b.log.Error("gagal mengupdate verifikasi otp", zap.Error(err))
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("failed: tidak terdapat usr yang membutuhkan otp")
	}
	return nil

}
