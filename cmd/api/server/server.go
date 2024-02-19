package server

import (
	"brick-task/internal/domain/payments/handler"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	HttpHandler *handler.Handler
	router      *echo.Echo
	Port        int
}

func (s *Server) Run() error {
	slog.Info(fmt.Sprintf("Starting HTTP server at :%d ...", s.Port))
	s.setupRouter()

	httpHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(s.router)

	srv := &http.Server{
		Handler:      httpHandler,
		Addr:         fmt.Sprintf(":%d", s.Port),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	return srv.ListenAndServe()
}
