package payments

import "context"

type UseCase interface {
	ValidateBankAccount(ctx context.Context, accountNumber string, bankCode string) (bool, error)
	Disburse(ctx context.Context, amount float64, accountNumber string, bankCode string) (bool, error)
}
