package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunStdin(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := Run(nil, strings.NewReader("# a\n\n## b\n"), &stdout, &stderr)

	if code != 0 {
		t.Fatalf("exit code = %d, stderr = %q", code, stderr.String())
	}
	if want := "[*** a]\n\n[** b]\n"; stdout.String() != want {
		t.Errorf("stdout = %q, want %q", stdout.String(), want)
	}
}

func TestRunNoInputWithoutPipe(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := Run(nil, nil, &stdout, &stderr)

	if code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
	if !strings.Contains(stderr.String(), "no input given") {
		t.Errorf("stderr = %q, want a \"no input given\" message", stderr.String())
	}
}

func TestRunFileInsertsTitle(t *testing.T) {
	path := filepath.Join(t.TempDir(), "note.md")
	if err := os.WriteFile(path, []byte("text\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer

	code := Run([]string{path}, nil, &stdout, &stderr)

	if code != 0 {
		t.Fatalf("exit code = %d, stderr = %q", code, stderr.String())
	}
	if want := "note\ntext\n"; stdout.String() != want {
		t.Errorf("stdout = %q, want %q", stdout.String(), want)
	}
}
