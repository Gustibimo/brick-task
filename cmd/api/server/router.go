package server

import "github.com/labstack/echo/v4"

func (s *Server) setupRouter() {
	s.Router = echo.New()
	v1 := s.Router.Group("/api/v1")
	v1.GET("/health", s.HttpHandler.HealthCheck)

	// Payment Routes
	payments := v1.Group("/payments")
	payments.POST("/validate-bank-account", s.HttpHandler.ValidateBankAccount)
	payments.POST("/webhook", s.HttpHandler.PaymentWebhook, s.Middleware.ValidateSignature)
	payments.POST("/disburse", s.HttpHandler.DoPayment)
}
