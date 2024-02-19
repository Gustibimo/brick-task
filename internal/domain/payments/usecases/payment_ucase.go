package usecases

import (
	"brick-task/internal/domain/payments"
	gateway "brick-task/internal/gateway/http"
)

type paymentUsecase struct {
	paymentRepo payments.Repository
	bankGateway gateway.BankClient
}

func NewPaymentUsecase(paymentRepo payments.Repository) payments.UseCase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
	}
}
