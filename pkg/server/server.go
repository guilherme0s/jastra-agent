package server

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilherme0s/atlans/pkg/config"
)

type Server struct {
	cfg        config.ServerSettings
	httpServer *http.Server
	router     *mux.Router
}

func New(cfg config.ServerSettings) *Server {
	cfg.SetDefaults()

	r := mux.NewRouter()

	return &Server{
		cfg:    cfg,
		router: r,
	}
}

func (s *Server) Start() error {
	addr := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.httpServer = &http.Server{
		Handler:      s.router,
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
	}

	slog.Info("Server is running on", "addr", listener.Addr().String())
	return s.httpServer.Serve(listener)
}

func (s *Server) Shutdown() {
	log.Print("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("%v", err)
	}
}
