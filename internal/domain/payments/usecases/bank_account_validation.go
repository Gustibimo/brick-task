package usecases

import "context"

func (u *paymentUsecase) ValidateBankAccount(ctx context.Context, accountNumber string, bankCode string) (bool, error) {
	// call API to validate bank account
	inquiryResponse, err := u.bankGateway.BankAccountInquiry(accountNumber, bankCode)
	if err != nil {
		return false, err
	}

	return inquiryResponse, nil
}
