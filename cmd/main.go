package main

import (
	"brick-task/cmd/api/server"
	"brick-task/db/postgres"
	"brick-task/internal/domain/payments/handler"
	"brick-task/internal/domain/payments/repositories"
	"brick-task/internal/domain/payments/usecases"
	"brick-task/pkg/pggen"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log/slog"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP API",
	Long:  "Run HTTP API for Brick Task Application",
	RunE:  runHttpCommand,
}

var (
	paymentHandler *handler.Handler
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	if err := httpCmd.Execute(); err != nil {
		slog.Error("error running server: ", err)
	}
}

func runHttpCommand(cmd *cobra.Command, args []string) error {
	initHTTP()

	httpServer := server.Server{
		HttpHandler: paymentHandler,
		Port:        8080,
	}

	return httpServer.Run()
}

func initHTTP() {
	postgres.Init()
	initInfra()
}

func initInfra() {
	dbConnection := pggen.Init("localhost", "5432", "postgres", "postgres", "brick_task", "brick")

	paymentRepository := repositories.NewPaymentRepository(dbConnection)
	paymentUsecase := usecases.NewPaymentUsecase(paymentRepository)

	paymentHandler = handler.NewHandler(paymentUsecase)

}
