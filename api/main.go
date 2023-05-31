package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/reki204/reki/config"
	"github.com/reki204/reki/server"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("failed to load config: %v", err)
		os.Exit(1)
	}

	srv := server.NewServer(cfg)

	// サーバーの起動
	go func() {
		if err := srv.Run(); err != nil {
			log.Printf("failed to run server: %v", err)
			os.Exit(1)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("failed to shut down server gracefully: %v", err)
		os.Exit(1)
	}

	log.Println("Server gracefully stopped")
}
