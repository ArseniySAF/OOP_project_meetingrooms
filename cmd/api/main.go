package main

import (
	"context"
	"errors"
	"log"
	"meeting-rooms/config"
	"meeting-rooms/internal/api"
	"meeting-rooms/internal/middleware"
	"meeting-rooms/internal/service"
	"meeting-rooms/internal/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.Load()
	_ = cfg

	repo, err := storage.NewPostgres()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer repo.Close()

	service := service.NewMeetingService(repo)
	handler := api.NewServer(service)

	r := chi.NewRouter()
	r.Use(middleware.JWTMiddleware)

	httpHandler := api.HandlerWithOptions(
		handler,
		api.ChiServerOptions{
			BaseRouter: r,
			BaseURL:    "/api/v1",
		},
	)

	server := &http.Server{
		Addr:    config.Load().Addr,
		Handler: httpHandler,
	}

	go func() {
		log.Printf("server started on port %s", config.Load().Port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("\nGraceful shutdown complete.")
}
