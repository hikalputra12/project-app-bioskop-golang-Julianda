package repository

import (
	"app-bioskop/internal/data/entity"
	"app-bioskop/pkg/database"
	"context"

	"go.uber.org/zap"
)

type PaymentMethodRepo struct {
	db  database.PgxIface
	log *zap.Logger
}

type PaymentMethodInterface interface {
	GetAllPaymentMethods(ctx context.Context) ([]*entity.PaymentMethod, error)
}

func NewPaymentMethod(db database.PgxIface, log *zap.Logger) PaymentMethodInterface {
	return &PaymentMethodRepo{
		db:  db,
		log: log,
	}
}

func (p *PaymentMethodRepo) GetAllPaymentMethods(ctx context.Context) ([]*entity.PaymentMethod, error) {
	query := `SELECT name,logo_url FROM payment_methods
ORDER BY id ASC `

	rows, err := p.db.Query(ctx, query)
	if err != nil {
		p.log.Error("failed to get all payment method on database", zap.Error(err))
	}
	var methods []*entity.PaymentMethod

	for rows.Next() {
		var t entity.PaymentMethod
		err := rows.Scan(&t.MethodName, &t.Logo)
		if err != nil {
			p.log.Error("failed to scan rows", zap.Error(err))
		}
		methods = append(methods, &t)
	}
	return methods, nil
}
