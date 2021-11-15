// Command gogenlicense finds and formats package licenses for use in go programs.
package main

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
			fmt.Println(legal.Notices)
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

	flag.StringVar(&flagDeclarationName, "n", flagDeclarationName, "Name of declaration to generate (default 'Notices'")
	defer func() {
		if flagDeclarationName != "" {
			return
		}
		flagDeclarationName = "Notices"
	}()

	flag.StringVar(&flagOutFile, "d", flagOutFile, "file to write output to (default '${basename($GOFILE)}_notices.go' or 'legal.go')")
	defer func() {
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

	flag.BoolVar(&flagModule, "m", flagModule, "exclude the path of the nearest module (folder with a 'go.mod' file")
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

var globalContext context.Context

func init() {
	var cancel context.CancelFunc
	globalContext, cancel = context.WithCancel(context.Background())

	cancelChan := make(chan os.Signal)
	signal.Notify(cancelChan, os.Interrupt)

	go func() {
		<-cancelChan
		cancel()
	}()
}
