package test

import (
	"os"
	"strings"
	"testing"

	"github.com/bhcoder23/gin-layout/internal/components"
	"github.com/bhcoder23/gin-layout/internal/utils"
	"github.com/spf13/viper"
)

func TestCommittedConfigsDoNotContainKnownSecrets(t *testing.T) {
	for _, path := range []string{"../configs/config.yaml", "../configs/config.example.yaml"} {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		text := string(data)
		for _, blocked := range []string{"123" + "456", "root:" + "123" + "456", `secret: "` + "key" + `"`} {
			if strings.Contains(text, blocked) {
				t.Fatalf("%s contains blocked secret-like value %q", path, blocked)
			}
		}
	}
}

func TestRedactDSNRemovesPassword(t *testing.T) {
	dsn := "root:secret@tcp(localhost:3306)/app?charset=utf8&parseTime=True"
	redacted := components.RedactDSN(dsn)

	if strings.Contains(redacted, "secret") {
		t.Fatalf("redacted DSN still contains password: %s", redacted)
	}
	if !strings.Contains(redacted, "root:***@tcp(localhost:3306)/app") {
		t.Fatalf("redacted DSN = %q", redacted)
	}
}

func TestComponentsInitReturnsConfigError(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	err := components.Init()
	if err == nil {
		t.Fatal("Init() error = nil, want missing config error")
	}
	if !strings.Contains(err.Error(), "mysql.dsn") {
		t.Fatalf("Init() error = %v, want mysql.dsn context", err)
	}
}

func TestGenerateTokenReturnsMissingSecretError(t *testing.T) {
	viper.Reset()
	defer viper.Reset()

	_, err := utils.GenerateToken(1)
	if err == nil {
		t.Fatal("GenerateToken() error = nil, want missing jwt.secret error")
	}
	if !strings.Contains(err.Error(), "jwt.secret") {
		t.Fatalf("GenerateToken() error = %v, want jwt.secret context", err)
	}
}
