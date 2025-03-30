// Package gogenlicense implements the gogenlicense command.
// See github.com/tkw1536/gogenlicense/cmd/gogenlicense.
package gogenlicense

// spellchecker: words stdlib libraryurl findgit gogenlicense

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/google/go-licenses/v2/licenses"
)

// Library is an internal representation of a library
type Library struct {
	Path  string
	Title string // like path, but with "Module" prefix

	LicenseName string
	LicenseText string
	LibraryURL  string
}

//go:embed resources/stdlib-license.txt
var goStdLibLicense string

// goStandardLibrary is the Go Standard Library
var goStandardLibrary = Library{
	Path:        "",
	Title:       "Go Standard Library",
	LicenseName: "BSD-3-Clause",
	LicenseText: goStdLibLicense,
	LibraryURL:  "https://golang.org/LICENSE",
}

// Some of the code below has been adapted from the 'save' and 'csv' commands of the github.com/google/go-licenses/licenses package.
// It implements the same functionality as the original commands, however it does not output to the command line
// but into a string that is further processed.
//
// For a copy of the license, see legal/notices.go.

// find finds licenses according to options
func find(ctx context.Context, options Options) ([]Library, error) {

	classifier, err := licenses.NewClassifier()
	if err != nil {
		return nil, fmt.Errorf("unable to create classifier: %w", err)
	}

	paths := make([]string, len(options.ModulePaths))
	for idx, path := range options.ModulePaths {
		paths[idx] = filepath.Join(path, "...")
	}

	libs, err := licenses.Libraries(ctx, classifier, false, nil, paths...)
	if err != nil {
		return nil, fmt.Errorf("unable to find libraries: %w", err)
	}

	libraries := make([]Library, 0, len(libs))

	for _, lib := range libs {
		name := lib.Name()

		// do not include the 'root' paths inside the output
		if !options.IncludeRoots && contains(name, options.ModulePaths) {
			continue
		}

		// do not include anything where the license file is a .go file.
		// That's a potential bug in "github.com/google/go-licenses/licenses".
		if strings.EqualFold(filepath.Ext(lib.LicenseFile), ".go") {
			continue
		}

		licenses := lib.Licenses
		if len(licenses) == 0 {
			return nil, fmt.Errorf("license classifier returned no licenses for %q", lib.LicenseFile)
		}

		licenseText, err := os.ReadFile(lib.LicenseFile)
		if err != nil {
			return nil, fmt.Errorf("unable to read license file %q: %w", lib.LicenseFile, err)
		}

		url, _ := goLicensesLibraryFileURL(lib, context.Background(), time.Minute, lib.LicenseFile)

		libraries = append(libraries, Library{
			Path:  name,
			Title: fmt.Sprintf("Module %s", name),

			LicenseName: licenses[0].Name,
			LibraryURL:  url,
			LicenseText: string(licenseText),
		})
	}

	// sort libraries
	sort.Slice(libraries, func(i, j int) bool {
		return libraries[i].Path > libraries[j].Path
	})

	return append([]Library{goStandardLibrary}, libraries...), nil
}

func contains(needle string, haystack []string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}
	return false
}
