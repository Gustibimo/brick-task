package repositories

import (
	"brick-task/internal/domain/payments/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(dbConnection *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{
		db: dbConnection,
	}
}

func (r *PaymentRepository) SavePayment(ctx context.Context, payment models.Payment) (*models.Payment, error) {
	var paymentResult models.Payment
	query := `INSERT INTO payments (uuid, external_id, amount, account_number, account_holder, bank_code, reference, description, status) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	_, err = r.db.ExecContext(ctx, query, payment.UUID, payment.ExternalID, payment.Amount, payment.AccountNumber, payment.AccountHolder, payment.BankCode, payment.Reference, payment.Description, payment.Status)
	err = tx.Commit()
	paymentResult = payment
	if err != nil {
		return nil, err
	}
	return &paymentResult, nil
}

func (r *PaymentRepository) UpdatePaymentStatus(ctx context.Context, status string, paymentId string) (bool, error) {
	query := `UPDATE payments SET status = $1 WHERE uuid = $2`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	_, err = r.db.ExecContext(ctx, query, status, paymentId)
	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}
