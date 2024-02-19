package gateway

import (
	"brick-task/internal/domain/payments/models"
	"net/http"
)

type BankClient interface {
	BankAccountInquiry(accountNumber string, bankCode string) (bool, error)
	DoPayment(paymentRequest models.DoPaymentRequest) (bool, error)
}

type bankClient struct {
}

func NewBankClient() BankClient {
	return &bankClient{}
}

func (b *bankClient) BankAccountInquiry(accountNumber string, bankCode string) (bool, error) {
	// call API to validate bank account
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.paystack.co/bank/resolve?account_number="+accountNumber+"&bank_code="+bankCode, nil)
	if err != nil {
		return false, err
	}

	response, err := client.Do(req)

	if err != nil {
		return false, err
	}

	if response.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

func (b *bankClient) DoPayment(paymentRequest models.DoPaymentRequest) (bool, error) {
	// call API to do payment
	return true, nil
}
