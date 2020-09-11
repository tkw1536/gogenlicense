package gogenlicense

import "context"

// Options are options for the Format function.
type Options struct {
	// ModulePaths are go import paths to generate license information for.
	ModulePaths []string

	// By default the root modules are not used.
	// When set to true, include them.
	IncludeRoots bool

	// ConfidenceThreshold is the confidence threshold for the license classifier.
	// The default is 0.9.
	ConfidenceThreshold float64

	// Name of the package to output
	OutPackage string

	// Name of the declaration
	DeclarationName string
}

// Format formats licenses into a valid go program according to Options.
func Format(ctx context.Context, options Options) (string, error) {
	libs, err := find(ctx, options)
	if err != nil {
		return "", err
	}
	result, err := generate(libs, options)
	if err != nil {
		return "", err
	}
	return gofmt(result)
}
