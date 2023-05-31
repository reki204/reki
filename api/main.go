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

// func main() {
// 	if err := run(context.Background()); err != nil {
// 		log.Printf("failed to terminated server: %v", err)
// 		os.Exit(1)
// 	}
// }

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	srv := server.NewServer(ctx, cfg)
	if err := srv.Run(); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

// func run(ctx context.Context) error {
// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		return err
// 	}

// 	// サーバー起動
// 	log.Printf("Listening and serving HTTP on :%v", cfg.Port)
// 	listenServer := server.Init(ctx)

// 	// Graceful Shutdown
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
// 	log.Println("Shutting down server...")

// 	if err := listenServer.Shutdown(ctx); err != nil {
// 		return err
// 	}

// 	log.Println("Server gracefully stopped")

// 	return nil
// }
