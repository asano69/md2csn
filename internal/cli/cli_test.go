package cli

import (
	"bytes"
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
