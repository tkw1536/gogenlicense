package gogenlicense

import (
	"context"
	"fmt"
	"go/format"
)

// spellchecker: words gofmt

// Options are options for the Format function.
type Options struct {
	// When set to true, ignore libraries with no license
	SkipNoLicense bool

	// ModulePaths are go import paths to generate license information for.
	ModulePaths []string

	// By default the root modules are not used.
	// When set to true, include them.
	IncludeRoots bool

	// Should test sources be included?
	IncludeTests bool

	// Name of the package to output
	OutPackage string

	// Name of the declaration
	DeclarationName string
}

// Format formats licenses according to options
func Format(ctx context.Context, options Options) (res string, err error) {
	resChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		libs, err := find(ctx, options)
		if err != nil {
			errChan <- err
			return
		}
		src, err := generate(libs, options)
		if err != nil {
			errChan <- err
			return
		}
		fmtSrc, err := gofmt(src)
		if err != nil {
			errChan <- err
			return
		}
		resChan <- fmtSrc
	}()

	select {
	case res = <-resChan:
		return
	case err = <-errChan:
		return
	case <-ctx.Done():
		return "", ctx.Err()
	}

}

// gofmt formats code in the standard go fmt style.
func gofmt(code string) (string, error) {
	bytes, err := format.Source([]byte(code))
	if err != nil {
		return "", fmt.Errorf("error running gofmt: %w", err)
	}
	return string(bytes), nil
}
