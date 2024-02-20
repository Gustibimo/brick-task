package models

import (
	"github.com/gofrs/uuid"
)

type AccountValidationRequest struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
}

type DoPaymentRequest struct {
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
	BankCode      string `json:"bank_code"`
	Amount        int    `json:"amount"`
	Reference     string `json:"reference"`
	Description   string `json:"description"`
	ExternalId    string `json:"external_id"`
	CallbackUrl   string `json:"callback_url"`
}

type PaymentStatusRequest struct {
	Status    string `json:"status"`
	PaymentId string `json:"payment_id"`
}

type Payment struct {
	UUID          uuid.UUID `json:"uuid"`
	ExternalID    string    `json:"external_id"`
	Amount        int       `json:"amount"`
	AccountNumber string    `json:"account_number"`
	AccountHolder string    `json:"account_holder"`
	BankCode      string    `json:"bank_code"`
	Reference     string    `json:"reference"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
}
