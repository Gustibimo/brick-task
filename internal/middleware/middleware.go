package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomMiddleware struct {
}

func NewCustomMiddleware() *CustomMiddleware {
	return &CustomMiddleware{}
}

func (m *CustomMiddleware) ValidateSignature(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		signatureKey := c.Request().Header.Get("X-Signature-Key")

		// simple logic for demonstration purposes
		// it should be a secret key + hashed request (HMAC) that is verified against the  signature key
		// Example: Verify if the signature key is valid
		if signatureKey != "webhook-123-xyz" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid signature key")
		}

		// Call the next handler in the chain
		return next(c)
	}
}
