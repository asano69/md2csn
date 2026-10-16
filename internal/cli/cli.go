// Package cli implements the md2csn command line interface.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/asano69/md2csn/internal/convert"
)

const usage = `Usage: md2csn [-o output] input...
       md2csn [-o output] < input.md

Convert Markdown files to Cosense notation.

  input   a Markdown file or a folder (searched recursively); may be repeated
  -o      output file or folder
          without -o, a single input file is written to stdout

Without input, Markdown is read from standard input and written to stdout.
`

type options struct {
	inputs []string
	output string
}

// Run runs the command and returns the process exit code.
// stdin must be nil when there is no piped input (for example, a terminal).
func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	opts, err := parseArgs(args)
	if errors.Is(err, flag.ErrHelp) {
		fmt.Fprint(stdout, usage)
		return 0
	}
	if err == nil {
		err = run(opts, stdin, stdout)
	}
	if err != nil {
		fmt.Fprintln(stderr, "md2csn:", err)
		return 1
	}
	return 0
}

// parseArgs accepts flags before, between, and after the inputs.
func parseArgs(args []string) (options, error) {
	var opts options
	fs := flag.NewFlagSet("md2csn", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // errors are reported by Run
	fs.StringVar(&opts.output, "o", "", "output file or folder")

	for {
		if err := fs.Parse(args); err != nil {
			return opts, err
		}
		if fs.NArg() == 0 {
			break
		}
		opts.inputs = append(opts.inputs, fs.Arg(0))
		args = fs.Args()[1:]
	}
	return opts, nil
}

// run converts every planned file. A failure does not stop the remaining files.
func run(opts options, stdin io.Reader, stdout io.Writer) error {
	if len(opts.inputs) == 0 {
		return convertStdin(opts.output, stdin, stdout)
	}
	jobs, err := planJobs(opts)
	if err != nil {
		return err
	}
	var errs []error
	for _, j := range jobs {
		if err := convertFile(j, stdout); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func convertFile(j job, stdout io.Writer) error {
	src, err := os.ReadFile(j.src)
	if err != nil {
		return fmt.Errorf("cannot read %s: %w", j.src, err)
	}
	return writeResult(j.dst, convert.Convert(src), stdout)
}

// convertStdin converts piped Markdown. Without -o the result goes to stdout.
func convertStdin(output string, stdin io.Reader, stdout io.Writer) error {
	if stdin == nil {
		return errors.New("no input given (run with -h for usage)")
	}
	src, err := io.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("cannot read standard input: %w", err)
	}
	return writeResult(output, convert.Convert(src), stdout)
}

// writeResult writes out to dst, or to stdout when dst is empty.
func writeResult(dst, out string, stdout io.Writer) error {
	if dst == "" {
		_, err := io.WriteString(stdout, out)
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("cannot create output folder for %s: %w", dst, err)
	}
	if err := os.WriteFile(dst, []byte(out), 0o644); err != nil {
		return fmt.Errorf("cannot write %s: %w", dst, err)
	}
	return nil
}
