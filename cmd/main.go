package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/omaraliali1010/go_template/api/route"
	"github.com/omaraliali1010/go_template/bootstrap"
)

func main() {

	// --- Bootstrap the app ---
	app := bootstrap.App()
	env := app.Env
	defer app.CloseDBConnection()

	// --- Setup context timeout ---
	timeout := time.Duration(env.ContextTimeout) * time.Second

	// --- Setup Chi router ---
	r := chi.NewRouter()

	// --- Middleware stack ---
	r.Use(middleware.RequestID) // Add request ID to each request
	r.Use(middleware.RealIP)    // Get real IP from X-Forwarded-For
	r.Use(middleware.Logger)    // Log requests
	r.Use(middleware.Recoverer) // Recover from panics
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin", "X-Requested-With"},
		ExposedHeaders:   []string{"Link", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	// --- Setup API routes ---
	// TODO: Create route
	route.Setup(app, timeout, r)

	// --- Create server ---
	srv := &http.Server{
		Addr:         env.ServerAddress,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// --- Run server in goroutine ---
	go func() {
		log.Printf("Starting server on %s", env.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// --- Graceful shutdown ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit // Wait for interrupt signal
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
