package payments

import (
	"brick-task/internal/domain/payments/models"
	"context"
)

type Repository interface {
	SavePayment(ctx context.Context, payment models.Payment) (*models.Payment, error)
	UpdatePaymentStatus(ctx context.Context, status string, paymentId string) (bool, error)
}
