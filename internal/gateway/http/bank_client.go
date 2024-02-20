package gateway

import (
	"brick-task/internal/domain/payments/models"
	"net/http"
)

type BankClient interface {
	BankAccountInquiry(accountNumber string, token string) (bool, error)
	DoPayment(paymentRequest models.DoPaymentRequest) (bool, error)
}

type bankClient struct {
}

func NewBankClient() BankClient {
	return &bankClient{}
}

func (b *bankClient) BankAccountInquiry(accountNumber string, bankCode string) (bool, error) {
	// call API to validate bank account
	bcaUrl := "https://65d34599522627d5010875bb.mockapi.io/api/v1/inquiry"
	mandiriUrl := "https://65d34599522627d5010875bb.mockapi.io/api/v1/inquiry"
	var url string
	if bankCode == "bca" {
		url = bcaUrl
	}

	if bankCode == "mandiri" {
		url = mandiriUrl
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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
	var url string

	bcaURL := "https://65d34599522627d5010875bb.mockapi.io/api/v1/doPayment"
	mandiriURL := "https://65d34599522627d5010875bb.mockapi.io/api/v1/doPayment"

	if paymentRequest.BankCode == "bca" {
		url = bcaURL
	}

	if paymentRequest.BankCode == "mandiri" {
		url = mandiriURL
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false, err
	}

	response, err := client.Do(req)

	if err != nil {
		return false, err
	}

	if response.StatusCode != http.StatusOK {
		return false, nil
	}

	if response.StatusCode == http.StatusOK {
		return true, nil
	}

	return true, nil
}
