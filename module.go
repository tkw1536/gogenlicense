package gogenlicense

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

var errGoModNotFound = errors.New("no 'go.mod' file found, unable to find current module")

// FindModuleRoot finds the directory containing the module at path.
// This works by finding a directory containing 'go.mod', starting recursively at path.
//
// Returns the path to a directory containing 'go.mod'.
// If no 'go.mod' file exists, returns an error.
func FindModuleRoot(path string) (root string, err error) {
	path = filepath.Clean(path)

	var prevPath string
	for path != prevPath {
		// check if the current path has a 'go.mod' file
		gomod := filepath.Join(path, "go.mod")
		if stat, err := os.Stat(gomod); err == nil && !stat.Mode().IsDir() {
			return path, nil
		}

		// go up to the next directory
		prevPath = path
		path = filepath.Dir(path)
	}
	return "", errGoModNotFound
}

// FindModulePath finds the package path to the module containing path.
// This works by finding the path to 'go.mod', starting recursively at path.
//
// Returns the package import path, if any.
// if no package import path exists, returns an error
func FindModulePath(path string) (pkg string, err error) {
	pkgroot, err := FindModuleRoot(path)
	if err != nil {
		return "", err
	}
	bytes, err := os.ReadFile(filepath.Join(pkgroot, "go.mod"))
	if err != nil {
		return "", err
	}
	f, err := modfile.ParseLax(path, bytes, nil)
	if err != nil {
		return "", err
	}
	return f.Module.Mod.Path, nil
}
