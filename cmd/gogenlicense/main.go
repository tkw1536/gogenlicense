// Command gogenlicense finds and formats package licenses for use in go programs.
// It is intended to comply with license requirements, but has not been vetted by a lawyer.
//
// Concretely this package generates go code that contains a single string variable which includes all legal notices.
// It also generates appropriate comments for godoc.
//
// # Quickstart
//
// gogenlicense is most commonly invoked using a 'go:generate' comment like so:
//
//	//go:generate gogenlicense -m
//
// After adding such a comment, run the 'go generate' command.
// Once a source file has been generated, your go code can make use of the generated notices.
// For example to add a legal flag that prints notices:
//
//	var legalFlag bool
//	flag.BoolVar(&legalFlag, "legal", legalFlag, "print legal notices and exit")
//	defer func() {
//		if legalFlag {
//			fmt.Println("This executable contains code from several different go packages. ")
//			fmt.Println("Some of these packages require licensing information to be made available to the end user. ")
//			fmt.Println(legal.Notices) // this references the generated constant!
//			os.Exit(0)
//		}
//	}()
//
// For a more complete example, see the source code of this command.
// The notices constant is implemented in the github.com/tkw1536/gogenlicense/legal subpackage.
// The legal flag is implemented in the main package.
//
// For usage of the go generate command, see https://golang.org/pkg/cmd/go/internal/generate/.
//
// # Usage
//
// The full invocation of the command is as follows:
//
//	gogenlicense [-legal] [-help] [-o FILE] [-t THRESHOLD] [-p PACKAGE] [-n NAME] [-m] [MODULE [MODULE...]]
//	-legal
//	  Print legal notices and exit.
//	-help
//	  Print a usage message and exit.
//	-p PACKAGE
//	  Package name of go source code file to generate.
//	  Defaults to the 'GOPACKAGE' environment variable, or to 'notices' if unset.
//	-n string
//	  Name of declaration to generate (default 'Notices')
//	-t THRESHOLD
//	 Threshold to use for the license classifier, between 0 and 1.
//	 Defaults to 0.9.
//	-o filename
//	  Path to write output to.
//	  Defaults to appending a suffix '_notices' to the source file pointed to by the 'GOFILE' environment variable.
//	  If 'GOFILE' is not provided, uses the filename 'notices.go'.
//	MODULE
//	  Import path of module to scan for module dependencies.
//	-m
//	  automatically find the current 'go module' containing the working directory and include it in MODULE list.
//
// # Acknowledgements
//
// This command was inspired by the https://classic.yarnpkg.com/lang/en/docs/cli/licenses/#toc-yarn-licenses-generate-disclaimer command.
// Internally this command relies on the excellent https://github.com/google/go-licenses package.
// See also the legal subpackage for full licenses of used packages.
package main

// spellchecker: words quickstart gopackage gofile glog gofmt subpackage gogenlicense godoc workdir licenseclassifier

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/tkw1536/gogenlicense"
	"github.com/tkw1536/gogenlicense/legal"
)

func main() {
	result, err := gogenlicense.Format(globalContext, gogenlicense.Options{
		ModulePaths:         argImportPaths,
		ConfidenceThreshold: flagConfidenceThreshold,

		OutPackage:      flagOutPackage,
		DeclarationName: flagDeclarationName,
	})
	if err != nil {
		log.Fatal(err)
	}

	if flagOutFile == "" {
		fmt.Println(result)
		return
	}

	err = os.WriteFile(flagOutFile, []byte(result), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

//
// Arguments & Defaults
//

var argImportPaths []string
var flagConfidenceThreshold float64
var flagOutFile string
var flagOutPackage string
var flagDeclarationName string
var flagModule bool

func init() {
	// HACK HACK HACK: remove glog flags by re-initializing the flagset
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var legalFlag bool
	flag.BoolVar(&legalFlag, "legal", legalFlag, "print legal notices and exit")
	defer func() {
		if legalFlag {
			fmt.Println("This executable contains code from several different go packages. ")
			fmt.Println("Some of these packages require licensing information to be made available to the end user. ")
			fmt.Print(legal.Notices)
			os.Exit(0)
		}
	}()

	flag.StringVar(&flagOutPackage, "p", flagOutPackage, "name of package to generate go file for (default '$GOPACKAGE' environment variable or 'legal')")
	defer func() {
		if flagOutPackage != "" {
			return
		}
		flagOutPackage = os.Getenv("GOPACKAGE")
		if flagOutPackage == "" {
			flagOutPackage = "legal"
		}
	}()

	flag.StringVar(&flagDeclarationName, "n", flagDeclarationName, "Name of declaration to generate (default 'Notices')")
	defer func() {
		if flagDeclarationName != "" {
			return
		}
		flagDeclarationName = "Notices"
	}()

	var legacyOutFile string
	flag.StringVar(&flagOutFile, "o", flagOutFile, "file to write output to (default '${basename($GOFILE)}_notices.go' or 'legal.go')")
	flag.StringVar(&legacyOutFile, "d", legacyOutFile, "legacy alias for -o")
	defer func() {
		if legacyOutFile != "" && flagOutFile == "" {
			flagOutFile = legacyOutFile
		}
		if flagOutFile != "" {
			return
		}
		envInFile := os.Getenv("GOFILE")
		if envInFile != "" {
			flagOutFile = strings.TrimSuffix(envInFile, filepath.Ext(envInFile)) + "_notices.go"
			return
		}
		flagOutFile = "legal.go"
	}()

	flagConfidenceThreshold = 0.9
	flag.Float64Var(&flagConfidenceThreshold, "t", flagConfidenceThreshold, "Threshold for the licenseclassifier")

	flag.BoolVar(&flagModule, "m", flagModule, "automatically use current 'go module' path as an import path")
	defer func() {
		if !flagModule {
			return
		}
		workdir, err := os.Getwd()
		if err != nil {
			glog.Fatal(err)
		}
		name, err := gogenlicense.FindModulePath(workdir)
		if err != nil {
			glog.Fatal(err)
		}
		argImportPaths = append(argImportPaths, name)
	}()

	flag.Parse()
	argImportPaths = append(argImportPaths, flag.Args()...)
}

//
// CTRL+C cancels a global context
//

var globalContext, _ = signal.NotifyContext(context.Background(), os.Interrupt)
