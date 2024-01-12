package transport

import (
	"archi/config"
	"archi/transport/handlers"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"time"
)

// transport/http/server.go

type Server struct {
	HTTP   *echo.Echo
	h      *handlers.Handlers
	config *config.Config
}

func NewServer(h *handlers.Handlers, config *config.Config) *Server {
	return &Server{h: h, config: config}
}

func (s *Server) Run(ctx context.Context) error {
	s.HTTP = s.BuildEngine()

	s.routes()

	go func() {
		err := s.HTTP.Start(s.config.Port)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutdown failed")
		}
	}()
	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()

	err := s.HTTP.Shutdown(ctxShutdown)
	if err != nil {
		log.Fatalf("Shutdown failed")
	}
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	return e
}
