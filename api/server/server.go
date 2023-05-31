package server

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reki204/reki/config"
	"github.com/reki204/reki/handler"
	routes "github.com/reki204/reki/router"
)

type Server struct {
	cfg        *config.ConfigList
	server     *http.Server
	serverCtx  context.Context
	serverDone context.CancelFunc
}

func NewServer(ctx context.Context, cfg *config.ConfigList) *Server {
	router := NewRouter(ctx)

	// Initialize routing
	if err := routes.SetRouting(ctx, router); err != nil {
		log.Fatalf(err.Error())
	}

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Port),
		Handler: router,
	}

	serverCtx, serverDone := context.WithCancel(ctx)

	return &Server{
		cfg:        cfg,
		server:     server,
		serverCtx:  serverCtx,
		serverDone: serverDone,
	}
}

func (s *Server) Run() error {
	log.Printf("Listening and serving HTTP on :%v", s.cfg.Port)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(err.Error())
		}
	}()

	// Graceful Shutdown
	<-s.serverCtx.Done()
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server gracefully stopped")

	return nil
}

func NewRouter(ctx context.Context) *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	r.Use(handler.SetCORS())

	return r
}

// Init initializes the server.
// func Init(ctx context.Context) *http.Server {
// 	router := NewRouter(ctx)
// 	// Initialize routing
// 	if err := routes.SetRouting(ctx, router); err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	srv := &http.Server{
// 		Addr:    ":8080",
// 		Handler: router,
// 	}

// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 		log.Fatalf(err.Error())
// 	}

// 	return srv
// }

// // NewRouter creates a new Gin router.
// func NewRouter(ctx context.Context) *gin.Engine {
// 	r := gin.Default()
// 	r.HandleMethodNotAllowed = true

// 	r.Use(handler.SetCORS())

// 	return r
// }
