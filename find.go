// Package gogenlicense implements the gogenlicense command.
// See github.com/tkw1536/gogenlicense/cmd/gogenlicense.
package gogenlicense

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/google/go-licenses/licenses"
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

// UnableToCreateClassifier indidicates an error creating a licenses classifier
type UnableToCreateClassifier struct {
	Err error
}

func (ua UnableToCreateClassifier) Error() string {
	return fmt.Sprintf("unable to create classifier: %s", ua.Err)
}

func (ua UnableToCreateClassifier) Unwrap() error {
	return ua.Err
}

// UnableToFindLibraries indidicates an error trying to find libraries
type UnableToFindLibraries struct {
	Err error
}

func (ul UnableToFindLibraries) Error() string {
	return fmt.Sprintf("unable to find libraries: %s", ul.Err)
}

func (ul UnableToFindLibraries) Unwrap() error {
	return ul.Err
}

// UnableToIdentifyLicense indicates an error attempting to identify the given license
type UnableToIdentifyLicense struct {
	Path string
	Err  error
}

func (ui UnableToIdentifyLicense) Error() string {
	return fmt.Sprintf("unable to identify license at %q: %s", ui.Path, ui.Err)
}

func (ui UnableToIdentifyLicense) Unwrap() error {
	return ui.Err
}

// UnableToReadLicense indicates an error attempting to read the given license
type UnableToReadLicense struct {
	Path string
	Err  error
}

func (ur UnableToReadLicense) Error() string {
	return fmt.Sprintf("unable to read license at %q: %s", ur.Path, ur.Err)
}

func (ur UnableToReadLicense) Unwrap() error {
	return ur.Err
}

// ErrUnknownLicense indicates that no known license is found in the given file
var ErrUnknownLicense = errors.New("unknown license")

// Some of the code below has been adapted from the 'save' and 'csv' commands of the github.com/google/go-licenses/licenses package.
// It implements the same functionality as the original commands, however it does not output to the command line
// but into a string that is further processed.
//
// For a copy of the license, see legal/notices.go.

// find finds licenses according to options
func find(ctx context.Context, options Options) ([]Library, error) {

	classifier, err := licenses.NewClassifier(options.ConfidenceThreshold)
	if err != nil {
		return nil, UnableToCreateClassifier{Err: err}
	}

	paths := make([]string, len(options.ModulePaths))
	for idx, path := range options.ModulePaths {
		paths[idx] = filepath.Join(path, "...")
	}

	libs, err := licenses.Libraries(ctx, classifier, nil, paths...)
	if err != nil {
		return nil, UnableToFindLibraries{Err: err}
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
		if strings.EqualFold(filepath.Ext(lib.LicensePath), ".go") {
			continue
		}

		licenseName, _, err := classifier.Identify(lib.LicensePath)
		if err != nil {
			return nil, UnableToIdentifyLicense{Path: lib.LicensePath, Err: err}
		}
		if licenseName == "" {
			return nil, UnableToIdentifyLicense{Path: lib.LicensePath, Err: ErrUnknownLicense}
		}
		licenseText, err := os.ReadFile(lib.LicensePath)
		if err != nil {
			return nil, UnableToReadLicense{Path: lib.LicensePath, Err: err}
		}

		libraries = append(libraries, Library{
			Path:  name,
			Title: fmt.Sprintf("Module %s", name),

			LicenseName: licenseName,
			LibraryURL:  libraryurl(ctx, lib),
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

var libraryGitRemotes = []string{
	"origin", "upstream",
}

// libraryurl returns the url to a library.
// If something goes wrong, silently returns the empty url
func libraryurl(ctx context.Context, lib *licenses.Library) string {
	repo, err := findgit(lib.LicensePath)
	if err != nil {
		// Can't find Git repo (possibly a Go Module?) - derive URL from lib name instead.
		url, err := lib.FileURL(ctx, lib.LicensePath)
		if err != nil {
			return ""
		}
		return url
	}

	for _, remote := range libraryGitRemotes {
		url, err := repo.FileURL(lib.LicensePath, remote)
		if err != nil {
			continue
		}
		return url.String()
	}

	return ""
}

var errNotFound = fmt.Errorf("no git repository that isn't also in parent")

// findgit finds the git repository at path, but excludes any that is potentially found in GOROOT.
func findgit(path string) (*licenses.GitRepo, error) {
	candidate, err := licenses.FindGitRepo(path)
	if err != nil {
		return nil, err
	}

	for _, repo := range forbidden {
		if *repo == *candidate {
			return nil, errNotFound
		}
	}

	return candidate, nil
}

// forbidden holds forbidden go repositories that should never be returned by findgit
var forbidden []*licenses.GitRepo

func init() {
	root, err := licenses.FindGitRepo(filepath.Join(runtime.GOROOT(), "_"))
	if err != nil {
		return
	}
	forbidden = append(forbidden, root)
}
