package handler

import (
	"brick-task/internal/domain/payments"
	"brick-task/internal/domain/payments/models"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Handler struct {
	PaymentsUseCase payments.UseCase
}

func NewHandler(paymentsUseCase payments.UseCase) *Handler {
	handler := &Handler{
		PaymentsUseCase: paymentsUseCase,
	}
	return handler
}

func (h *Handler) ValidateBankAccount(c echo.Context) error {

	var accountValidationRequest models.AccountValidationRequest

	if err := c.Bind(&accountValidationRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	slog.Info("Validating bank account: ", accountValidationRequest.AccountNumber, accountValidationRequest.BankCode)
	if accountValidationRequest.AccountNumber == "" || accountValidationRequest.BankCode == "" {
		return c.JSON(http.StatusBadRequest, BaseResponse{
			Status:  "invalid request",
			Message: "Account number and bank code are required",
		})
	}

	valid, err := h.PaymentsUseCase.ValidateBankAccount(c.Request().Context(), accountValidationRequest.AccountNumber,
		accountValidationRequest.BankCode)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"valid": valid,
	})
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"status": "ok",
	})
}

type BaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
