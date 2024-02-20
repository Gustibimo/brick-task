package payments

import (
	"brick-task/internal/domain/payments/models"
	"context"
)

type UseCase interface {
	ValidateBankAccount(ctx context.Context, accountNumber string, bankCode string) (bool, error)
	Disburse(ctx context.Context, request models.DoPaymentRequest) (*models.Payment, error)
	UpdatePaymentStatus(ctx context.Context, request models.PaymentStatusRequest) (bool, error)
}
