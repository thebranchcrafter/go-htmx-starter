package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thebranchcrafter/go-htmx-starter/internal/app"
	"github.com/thebranchcrafter/go-htmx-starter/internal/app/config"
)

func main() {
	// Load the configuration from the .env file
	cnf := config.LoadConfig()

	k, err := app.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}

	// Start the server in a goroutine
	serverErrCh := make(chan error)
	go func() {
		log.Println("Starting server...")
		if err := k.StartServer(cnf.AddressPort); err != nil {
			serverErrCh <- err
		}
	}()

	// Gracefully handle shutdown signals (SIGINT, SIGTERM)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrCh:
		// If server fails to start or crashes
		log.Fatalf("Server failed: %v", err)
	case sig := <-signalCh:
		// If we receive an OS signal
		log.Printf("Received signal: %s. Shutting down server...", sig)
	}

	// Graceful shutdown of the server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := k.ShutdownServer(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}
