// Package test contains black-box checks for the application scaffold.
package test

import (
	"os"
	"path/filepath"
	"testing"

	appconfig "github.com/bhcoder23/gin-layout/internal/config"
	"github.com/spf13/viper"
)

func TestConfigLoadDefaultsToConfigYAML(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	dir := t.TempDir()
	configDir := filepath.Join(dir, "configs")
	if err := os.Mkdir(configDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte("server:\n  port: 18080\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatal(err)
		}
	}()

	if err := appconfig.Load("", ""); err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if got := filepath.Base(viper.ConfigFileUsed()); got != "config.yaml" {
		t.Fatalf("ConfigFileUsed() = %q, want config.yaml", got)
	}
	if got := viper.GetInt("server.port"); got != 18080 {
		t.Fatalf("server.port = %d, want 18080", got)
	}
}
