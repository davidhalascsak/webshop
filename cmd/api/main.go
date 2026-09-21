package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"webshop/internal/auth"
	"webshop/internal/config"
	"webshop/internal/product"
	"webshop/internal/router"
)

func main() {
	cfg := config.Load()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	port := cfg.Port

	db, err := sqlx.Connect("postgres", cfg.DatabaseUrl)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Fatal("failed to ping database", zap.Error(err))
	}
	logger.Info("database connected successfully")

	authenticator, err := auth.NewAuthenticator(
		context.Background(),
		cfg.KeycloakIssuerURL,
		cfg.KeycloakClientID,
	)
	if err != nil {
		logger.Fatal("failed to initialize authenticator", zap.Error(err))
	}

	productDatabase := product.ProductRepository(db)
	productService := product.ProductService(productDatabase)
	productHandler := product.ProductHandler(productService)
	handlers := router.Handlers{
		ProductHandler: productHandler,
	}

	engine := router.Setup(handlers, authenticator)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Info("starting server", zap.String("port", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("forced shutdown", zap.Error(err))
	}
}
