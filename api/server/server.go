package server

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reki204/reki/config"
	"github.com/reki204/reki/middleware"
	routes "github.com/reki204/reki/router"
)

type Server struct {
	cfg    *config.ConfigList
	router *gin.Engine
	server *http.Server
}

func NewServer(cfg *config.ConfigList) *Server {
	// ルーターの初期化
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.Use(middleware.SetCORS())
	routes.SetRouting(router)

	// サーバーの設定
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Port),
		Handler: router,
	}

	return &Server{
		cfg:    cfg,
		router: router,
		server: server,
	}
}

func (s *Server) Run() error {
	log.Printf("Listening and serving HTTP on :%v", s.cfg.Port)
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
