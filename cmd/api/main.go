
package main

import (
 "log"
 "meeting-rooms/config"
 "meeting-rooms/internal/api"
 "meeting-rooms/internal/middleware"
 "meeting-rooms/internal/service"
 "meeting-rooms/internal/storage"
 "net/http"

 "github.com/go-chi/chi/v5"
)

func main() {
 repo, err := storage.NewPostgres()
 if err != nil {
  log.Fatalf("failed to connect to postgres: %v", err)
 }
 defer repo.Close()

 service := service.NewMeetingService(repo)
 handler := api.NewServer(service)

 r := chi.NewRouter()
 r.Use(middleware.JWTMiddleware)

 httpHandler := api.HandlerFromMux(handler, r)

 server := &http.Server{
  Addr:    config.Load().Addr,
  Handler: httpHandler,
 }

 _ = server
}
