package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/max1t1a/subscription-service/config"
	_ "github.com/max1t1a/subscription-service/docs"
	"github.com/max1t1a/subscription-service/internal/api"
	paymentHandler "github.com/max1t1a/subscription-service/internal/api/handler/payment"
	subscriptionHandler "github.com/max1t1a/subscription-service/internal/api/handler/subscription"
	paymentRepository "github.com/max1t1a/subscription-service/internal/repository/payment"
	subscriptionRepository "github.com/max1t1a/subscription-service/internal/repository/subscription"
	paymentService "github.com/max1t1a/subscription-service/internal/service/payment"
	subscriptionService "github.com/max1t1a/subscription-service/internal/service/subscription"
	"github.com/max1t1a/subscription-service/internal/worker"
)

// @title           Subscription Service API
// @version         1.0
// @description     REST API for managing user subscriptions and payments.
// @BasePath        /api/v1
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg := config.Load()

	db, err := connectDB(cfg.DB)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	subRepository := subscriptionRepository.New(db)
	payRepository := paymentRepository.New(db)

	subService := subscriptionService.New(subRepository)
	payService := paymentService.New(payRepository)

	subHandler := subscriptionHandler.New(subService)
	payHandler := paymentHandler.New(payService)

	router := api.NewRouter(subHandler, payHandler, logger)

	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	renewalWorker := worker.NewRenewalWorker(subRepository, payRepository, cfg.Worker.Interval, cfg.Worker.RenewalThreshold, logger)
	go renewalWorker.Start(ctx)

	go func() {
		logger.Info("server starting", zap.String("port", cfg.AppPort))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down...")
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown", zap.Error(err))
	}

	logger.Info("server stopped")
}

func connectDB(cfg config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
