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

Convert Markdown files to Cosense notation.

  input   a Markdown file or a folder (searched recursively); may be repeated
  -o      output file or folder
          without -o, a single input file is written to stdout
`

type options struct {
	inputs []string
	output string
}

// Run runs the command and returns the process exit code.
func Run(args []string, stdout, stderr io.Writer) int {
	opts, err := parseArgs(args)
	if errors.Is(err, flag.ErrHelp) {
		fmt.Fprint(stdout, usage)
		return 0
	}
	if err == nil {
		err = run(opts, stdout)
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
	if len(opts.inputs) == 0 {
		return opts, errors.New("no input given (run with -h for usage)")
	}
	return opts, nil
}

// run converts every planned file. A failure does not stop the remaining files.
func run(opts options, stdout io.Writer) error {
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
	out := convert.Convert(src)

	if j.dst == "" {
		_, err := io.WriteString(stdout, out)
		return err
	}
	if err := os.MkdirAll(filepath.Dir(j.dst), 0o755); err != nil {
		return fmt.Errorf("cannot create output folder for %s: %w", j.dst, err)
	}
	if err := os.WriteFile(j.dst, []byte(out), 0o644); err != nil {
		return fmt.Errorf("cannot write %s: %w", j.dst, err)
	}
	return nil
}
