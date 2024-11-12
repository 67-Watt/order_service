package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/http-swagger" // Swagger handler

	"order_service/config"
	"order_service/endpoint"
	"order_service/models" // Import models package for migration
	"order_service/repository"
	"order_service/transport"
	"order_service/usecase"
	"order_service/utils"
)

func main() {
	// Setup structured logger
	appLogger := utils.NewLogger()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := utils.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		dbConn, _ := db.DB()
		err := dbConn.Close()
		if err != nil {
			return
		}
	}()

	// Run migrations
	if err := db.AutoMigrate(&models.Order{}, &models.OrderDetail{}, &models.Customer{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Setup Prometheus metrics counter
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "order_usecase_requests_total",
		Help: "Total number of requests to the order use case",
	}, []string{"method"})
	prometheus.MustRegister(counter)

	// Initialize the repository and use case
	orderRepo := repository.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)

	// Wrap the use case with logging and metrics middleware
	orderUseCase = utils.LoggingMiddleware(appLogger)(orderUseCase)
	orderUseCase = utils.MetricsMiddleware(counter)(orderUseCase)

	// Create Endpoints and HTTP Handlers
	endpoints := endpoint.MakeEndpoints(orderUseCase)
	router := mux.NewRouter()
	router = transport.NewHTTPHandler(endpoints, router, appLogger)

	// Expose the Prometheus metrics endpoint and Swagger documentation
	router.Handle("/metrics", promhttp.Handler())
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Create and start HTTP server
	server := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Channel to listen for OS interrupt signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting server at %s\n", cfg.ServerAddress)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed: %v\n", err)
		}
	}()

	<-interrupt // Wait for the interrupt signal

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
