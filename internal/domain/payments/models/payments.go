package models

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
}
