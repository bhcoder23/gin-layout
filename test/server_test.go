package test

import (
	"net/http"
	"testing"
	"time"

	appserver "github.com/bhcoder23/gin-layout/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TestNewHTTPServerUsesConfiguredAddressAndTimeouts(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	viper.Set("server.readtimeout", 7)
	viper.Set("server.writetimeout", 11)

	srv := appserver.NewHTTPServer("127.0.0.1", 18082, gin.New())

	if srv.Addr != "127.0.0.1:18082" {
		t.Fatalf("Addr = %q, want 127.0.0.1:18082", srv.Addr)
	}
	if srv.ReadTimeout != 7*time.Second {
		t.Fatalf("ReadTimeout = %s, want 7s", srv.ReadTimeout)
	}
	if srv.WriteTimeout != 11*time.Second {
		t.Fatalf("WriteTimeout = %s, want 11s", srv.WriteTimeout)
	}
	if srv.MaxHeaderBytes != http.DefaultMaxHeaderBytes {
		t.Fatalf("MaxHeaderBytes = %d, want %d", srv.MaxHeaderBytes, http.DefaultMaxHeaderBytes)
	}
}

func TestNormalizeGinModeMapsRuntimeModes(t *testing.T) {
	tests := map[string]string{
		"":        gin.DebugMode,
		"debug":   gin.DebugMode,
		"dev":     gin.DebugMode,
		"release": gin.ReleaseMode,
		"prod":    gin.ReleaseMode,
		"test":    gin.TestMode,
	}

	for input, want := range tests {
		if got := appserver.NormalizeGinMode(input); got != want {
			t.Fatalf("NormalizeGinMode(%q) = %q, want %q", input, got, want)
		}
	}
}
