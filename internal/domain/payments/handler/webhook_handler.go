package handler

import (
	"brick-task/internal/domain/payments/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) PaymentWebhook(c echo.Context) error {

	// call update payment status for pending payment
	var paymentStatusRequest models.PaymentStatusRequest

	if err := c.Bind(&paymentStatusRequest); err != nil {
		return c.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if paymentStatusRequest.PaymentId == "" || paymentStatusRequest.Status == "" {
		return c.JSON(400, map[string]interface{}{
			"error": "payment_id and status are required",
		})
	}

	// update payment status
	_, err := h.PaymentsUseCase.UpdatePaymentStatus(c.Request().Context(), paymentStatusRequest)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "ok",
	})
}
