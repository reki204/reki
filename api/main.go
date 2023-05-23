package main

import (
	"context"
	"log"
	"os"

	"github.com/reki204/reki/config"
	"github.com/reki204/reki/server"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// サーバー起動
	log.Printf("Listening and serving HTTP on :%v", cfg.Port)
	listenServer := server.Init(ctx)
	return listenServer
}
