// Package server owns HTTP server startup and graceful shutdown.
package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bhcoder23/gin-layout/internal/components"
	"github.com/bhcoder23/gin-layout/internal/migrates"
	"github.com/bhcoder23/gin-layout/internal/routers"
	"github.com/bhcoder23/gin-layout/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// StartHTTPServer initializes dependencies, starts the HTTP server, and shuts it down gracefully.
func StartHTTPServer(ctx context.Context, host string, port int) error {
	if ctx == nil {
		ctx = context.Background()
	}

	gin.SetMode(NormalizeGinMode(viper.GetString("server.mode")))
	if err := components.Init(); err != nil {
		return err
	}
	if err := migrates.DoMigrate(); err != nil {
		return err
	}
	services.InitServices()
	r := routers.InitRouter()

	s := NewHTTPServer(host, port, r)
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	serverErr := make(chan error, 1)
	go func() {
		zap.L().Info("HTTP server started", zap.String("addr", s.Addr))
		if err := s.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	services.SetReady()
	defer services.SetNotReady()

	signalCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-serverErr:
		return err
	case <-signalCtx.Done():
	}

	zap.L().Warn("shutdown signal received")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), configDuration("server.shutdowntimeout", 5*time.Second))
	defer cancel()

	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(shutdownCtx); err != nil {
		return err
	}
	zap.L().Info("HTTP server stopped")
	return nil
}

// NewHTTPServer builds a configured HTTP server without starting network IO.
func NewHTTPServer(host string, port int, handler http.Handler) *http.Server {
	if port <= 0 {
		port = 8080
	}
	return &http.Server{
		Addr:           net.JoinHostPort(host, strconv.Itoa(port)),
		Handler:        handler,
		ReadTimeout:    configDuration("server.readtimeout", 0),
		WriteTimeout:   configDuration("server.writetimeout", 0),
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}
}

// NormalizeGinMode maps runtime aliases to Gin modes.
func NormalizeGinMode(mode string) string {
	switch strings.ToLower(mode) {
	case "release", "prod", "production":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}

func configDuration(key string, fallback time.Duration) time.Duration {
	seconds := viper.GetInt(key)
	if seconds <= 0 {
		return fallback
	}
	return time.Duration(seconds) * time.Second
}
