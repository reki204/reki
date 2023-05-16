package server

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/reki204/reki/handler"
	routes "github.com/reki204/reki/router"
)

// Init initializes the server.
func Init(ctx context.Context) error {
	router := NewRouter(ctx)
	// Initialize routing
	if err := routes.SetRouting(ctx, router); err != nil {
		return err
	}

	if err := router.Run(); err != nil {
		log.Fatalf(err.Error())
	}

	return nil
}

// NewRouter creates a new Gin router.
func NewRouter(ctx context.Context) *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	r.Use(handler.SetCORS(r))

	return r
}
