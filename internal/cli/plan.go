package cli

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// job is one file conversion. An empty dst means standard output.
type job struct {
	src string
	dst string
}

// input is a Markdown file found from the command line arguments.
type input struct {
	path   string
	rel    string // output path relative to the output folder
	direct bool   // given as a file argument, not found inside a folder
}

// planJobs decides which files to convert and where to write each result.
func planJobs(opts options) ([]job, error) {
	inputs, err := collectInputs(opts.inputs)
	if err != nil {
		return nil, err
	}
	if len(inputs) == 0 {
		return nil, errors.New("no Markdown files found in the given inputs")
	}

	single := len(inputs) == 1 && inputs[0].direct
	switch {
	case opts.output == "" && single:
		return []job{{src: inputs[0].path}}, nil
	case opts.output == "":
		return nil, errors.New("-o is required when converting a folder or multiple files")
	case single && !isFolderPath(opts.output):
		return []job{{src: inputs[0].path, dst: opts.output}}, nil
	default:
		return jobsInFolder(inputs, opts.output)
	}
}

func jobsInFolder(inputs []input, folder string) ([]job, error) {
	seen := make(map[string]string, len(inputs))
	jobs := make([]job, 0, len(inputs))
	for _, in := range inputs {
		dst := filepath.Join(folder, withTxtExt(in.rel))
		if prev, dup := seen[dst]; dup {
			return nil, fmt.Errorf("%s and %s would both be written to %s", prev, in.path, dst)
		}
		seen[dst] = in.path
		jobs = append(jobs, job{src: in.path, dst: dst})
	}
	return jobs, nil
}

func collectInputs(paths []string) ([]input, error) {
	var inputs []input
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			return nil, fmt.Errorf("cannot access input: %w", err)
		}
		if !info.IsDir() {
			inputs = append(inputs, input{path: p, rel: filepath.Base(p), direct: true})
			continue
		}
		found, err := walkMarkdown(p)
		if err != nil {
			return nil, fmt.Errorf("cannot search folder %s: %w", p, err)
		}
		inputs = append(inputs, found...)
	}
	return inputs, nil
}

// walkMarkdown finds Markdown files under root, recursively.
func walkMarkdown(root string) ([]input, error) {
	var found []input
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !isMarkdown(path) {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		found = append(found, input{path: path, rel: rel})
		return nil
	})
	return found, err
}

func isMarkdown(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".md", ".markdown":
		return true
	}
	return false
}

// isFolderPath reports whether p is an existing folder or is written with a trailing separator.
func isFolderPath(p string) bool {
	if strings.HasSuffix(p, "/") || strings.HasSuffix(p, string(filepath.Separator)) {
		return true
	}
	info, err := os.Stat(p)
	return err == nil && info.IsDir()
}

// titleFromPath returns the file name of p without its extension.
func titleFromPath(p string) string {
	base := filepath.Base(p)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func withTxtExt(p string) string {
	return strings.TrimSuffix(p, filepath.Ext(p)) + ".txt"
}
