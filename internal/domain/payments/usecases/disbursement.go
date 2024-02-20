package usecases

import (
	"brick-task/internal/domain/payments/models"
	"context"
	"github.com/gofrs/uuid"
)

func (u *paymentUsecase) Disburse(ctx context.Context, request models.DoPaymentRequest) (*models.Payment, error) {
	_, err := u.bankGateway.DoPayment(request)
	if err != nil {
		return nil, err
	}

	paymentEntity := models.Payment{
		UUID:          uuid.Must(uuid.NewV4()),
		ExternalID:    request.ExternalId,
		Amount:        request.Amount,
		AccountNumber: request.AccountNumber,
		AccountHolder: request.AccountHolder,
		BankCode:      request.BankCode,
		Reference:     request.Reference,
		Description:   request.Description,
		Status:        "pending",
	}

	paymentResult, err := u.paymentRepo.SavePayment(ctx, paymentEntity)
	if err != nil {
		return nil, err
	}

	return paymentResult, nil
}

func (u *paymentUsecase) UpdatePaymentStatus(ctx context.Context, request models.PaymentStatusRequest) (bool, error) {
	status, err := u.paymentRepo.UpdatePaymentStatus(ctx, request.Status, request.PaymentId)
	if err != nil {
		return false, err
	}

	return status, nil
}
