package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestServerHelpShowsConfigFlags(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "server", "--help")
	cmd.Dir = ".."

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go run . server --help failed: %v\n%s", err, output)
	}

	text := string(output)
	for _, want := range []string{"-c, --config string", "--env string", "gin-layout server -c configs/config.yaml"} {
		if !strings.Contains(text, want) {
			t.Fatalf("help output missing %q:\n%s", want, text)
		}
	}
}
