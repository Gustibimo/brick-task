package usecases

import "context"

func (u *paymentUsecase) Disburse(ctx context.Context, amount float64, accountNumber string, bankCode string) (bool, error) {
	panic("not implemented")
}
