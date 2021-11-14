package gogenlicense

import (
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

const goModName = "go.mod"

var errNoGoMod = errors.New("no go.mod")

// FindGoMod finds the 'go.mod' file for a given path
func FindGoMod(path string) (gomod string, err error) {
	var pastPath string
	for path != pastPath {
		// if gomod exists, return it
		gomod = filepath.Join(path, goModName)
		if stat, err := os.Stat(gomod); err == nil && stat.Mode().IsRegular() {
			return gomod, nil
		}

		// update the path
		pastPath = path
		path = filepath.Join(path, "..")
	}
	return "", errNoGoMod
}

// FindModulePath finds the path of the current module
func FindModulePath(path string) (name string, err error) {
	gomod, err := FindGoMod(path)
	if err != nil {
		return "", err
	}
	bytes, err := os.ReadFile(gomod)
	if err != nil {
		return "", err
	}
	f, err := modfile.ParseLax(path, bytes, nil)
	if err != nil {
		return "", err
	}
	return f.Module.Mod.Path, nil
}
