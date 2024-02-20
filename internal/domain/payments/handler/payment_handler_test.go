package handler

import (
	"brick-task/internal/domain/payments/mocks"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var paymentUCase = new(mocks.Usecase)

func TestHandler_ValidateBankAccount(t *testing.T) {
	handler := NewHandler(paymentUCase) // Replace nil with your mock payments use case for testing

	testCases := []struct {
		Name           string
		RequestBody    interface{}
		ExpectedStatus int
	}{
		{
			Name:           "ValidRequest",
			RequestBody:    map[string]string{"account_number": "1234567890", "bank_code": "123"},
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "InvalidRequest",
			RequestBody:    map[string]string{"account_number": "", "bank_code": ""},
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tc.RequestBody)
			req := httptest.NewRequest(http.MethodPost, "/validate-bank-account", bytes.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, rec)
			err := handler.ValidateBankAccount(ctx)

			paymentUCase.On("ValidateBankAccount", ctx.Request().Context(), "1234567890", "123").Return(true, nil)

			assert.NoError(t, err)
			assert.Equal(t, tc.ExpectedStatus, rec.Code)
		})
	}
}

func TestHandler_DoPayment(t *testing.T) {
	handler := NewHandler(paymentUCase) // Replace nil with your mock payments use case for testing

	testCases := []struct {
		Name           string
		RequestBody    interface{}
		ExpectedStatus int
	}{
		{
			Name:           "ValidRequest",
			RequestBody:    map[string]interface{}{"account_number": "1234567890", "bank_code": "123", "amount": 100},
			ExpectedStatus: http.StatusOK,
		},
		{
			Name:           "InvalidRequest",
			RequestBody:    map[string]interface{}{"account_number": "", "bank_code": "", "amount": 0},
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tc.RequestBody)
			req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, rec)
			err := handler.DoPayment(ctx)
			assert.NoError(t, err)
			assert.Equal(t, tc.ExpectedStatus, rec.Code)
		})
	}
}
