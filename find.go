// Package gogenlicense implements the gogenlicense command.
// See github.com/tkw1536/gogenlicense/cmd/gogenlicense.
package gogenlicense

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/google/go-licenses/licenses"
	"github.com/pkg/errors"
)

// Some of the code below has been adapted from the 'save' and 'csv' commands of the github.com/google/go-licenses/licenses package.
// It implements the same functionality as the original commands, however it does not output to the command line
// but into a string that is further processed.
//
// For a copy of the license, see legal/notices.go.

// find finds licenses according to options
func find(ctx context.Context, options Options) ([]Library, error) {

	classifier, err := licenses.NewClassifier(options.ConfidenceThreshold)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create classifier")
	}

	paths := make([]string, len(options.ModulePaths))
	for idx, path := range options.ModulePaths {
		paths[idx] = filepath.Join(path, "...")
	}

	libs, err := licenses.Libraries(ctx, classifier, paths...)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to find libraries")
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
			return nil, errors.Wrapf(err, "Unable to identify license %q", lib.LicensePath)
		}
		licenseText, err := os.ReadFile(lib.LicensePath)
		if err != nil {
			return nil, errors.Wrapf(err, "Unable to read license file %q", lib.LicensePath)
		}

		libraries = append(libraries, Library{
			Path:  name,
			Title: fmt.Sprintf("Module %s", name),

			LicenseName: licenseName,
			LibraryURL:  libraryurl(lib),
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
func libraryurl(lib *licenses.Library) string {
	repo, err := licenses.FindGitRepo(lib.LicensePath)
	if err != nil { // Can't find Git repo (possibly a Go Module?) - derive URL from lib name instead.
		url, err := lib.FileURL(lib.LicensePath)
		if err != nil {
			return ""
		}
		return url.String()
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

// Library is an internal representation of a library
type Library struct {
	Path  string
	Title string // like path, but with "Module" prefix

	LicenseName string
	LicenseText string
	LibraryURL  string
}

// goStandardLibrary is the Go Standard Library
var goStandardLibrary = Library{
	Path:        "",
	Title:       "Go Standard Library",
	LicenseName: "BSD-3-Clause",
	LicenseText: `Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

	* Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
	* Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
	* Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`,
	LibraryURL: "https://golang.org/LICENSE",
}
