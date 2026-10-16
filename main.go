package main

import (
	"io"
	"os"

	"github.com/asano69/md2csn/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], pipedStdin(), os.Stdout, os.Stderr))
}

// pipedStdin returns stdin when it is a pipe or a redirected file,
// and nil when it is a terminal.
func pipedStdin() io.Reader {
	info, err := os.Stdin.Stat()
	if err != nil || info.Mode()&os.ModeCharDevice != 0 {
		return nil
	}
	return os.Stdin
}
