package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CServer struct {
	port   int    `yaml:"port"`
	domain string `yaml:"domain"`
}

type Server struct {
	log    *logrus.Logger
	server *http.Server
}

func NewServer(cfg *CServer, log *logrus.Logger) *Server {
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.domain, cfg.port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!\n"))
		}),
	}
	return &Server{
		log:    log,
		server: server,
	}
}

func (s *Server) start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		s.log.Fatalf("Ошибка запуска сервера: %v", err)
		return err
	}
	return nil
}

func (s *Server) refresh() error {
	s.log.Info("Остановка сервера...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	// Ожидание остановки сервера
	for {
		s.log.Info("Ожидание остановки сервера...")
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			s.log.Errorf("Error waiting for server to stop: %s\n", err)
			time.Sleep(1 * time.Second)
			return err
		} else {
			s.log.Info("Сервер успешно остановлен")
			break
		}
	}
	return nil
}
