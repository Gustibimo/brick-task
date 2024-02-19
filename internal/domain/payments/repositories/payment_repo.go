package repositories

import (
	"github.com/jmoiron/sqlx"
)

type PaymentRepository struct {
	DbConnection *sqlx.DB
}

func NewPaymentRepository(dbConnection *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{
		DbConnection: dbConnection,
	}
}
