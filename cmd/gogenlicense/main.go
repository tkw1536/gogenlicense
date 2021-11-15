// Command gogenlicense finds and formats package licenses for use in go programs.
package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
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
		glog.Fatal(err)
	}

	if flagOutFile == "" {
		fmt.Println(result)
		return
	}

	err = ioutil.WriteFile(flagOutFile, []byte(result), os.ModePerm)
	if err != nil {
		glog.Fatal(err)
	}
}

//
// Arguments & Defaults
//

var flagset = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

var argImportPaths []string
var flagConfidenceThreshold float64
var flagOutFile string
var flagOutPackage string
var flagDeclarationName string
var flagModule bool

func init() {
	var legalFlag bool
	flagset.BoolVar(&legalFlag, "legal", legalFlag, "print legal notices and exit")
	defer func() {
		if legalFlag {
			fmt.Println("This executable contains code from several different go packages. ")
			fmt.Println("Some of these packages require licensing information to be made available to the end user. ")
			fmt.Println(legal.Notices)
			os.Exit(0)
		}
	}()

	flagset.StringVar(&flagOutPackage, "p", flagOutPackage, "name of package to generate go file for (default '$GOPACKAGE' environment variable or 'legal')")
	defer func() {
		if flagOutPackage != "" {
			return
		}
		flagOutPackage = os.Getenv("GOPACKAGE")
		if flagOutPackage == "" {
			flagOutPackage = "legal"
		}
	}()

	flagset.StringVar(&flagDeclarationName, "n", flagDeclarationName, "Name of declaration to generate (default 'Notices'")
	defer func() {
		if flagDeclarationName != "" {
			return
		}
		flagDeclarationName = "Notices"
	}()

	flagset.StringVar(&flagOutFile, "d", flagOutFile, "file to write output to (default '${basename($GOFILE)}_notices.go' or 'legal.go')")
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
	flagset.Float64Var(&flagConfidenceThreshold, "t", flagConfidenceThreshold, "Threshold for the licenseclassifier")

	flagset.BoolVar(&flagModule, "m", flagModule, "exclude the path of the nearest module (folder with a 'go.mod' file")
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

	flagset.Parse(os.Args[1:])
	argImportPaths = append(argImportPaths, flag.Args()...)
}

func init() {
	// HACK HACK HACK
	// we need this supress various glog warnings
	flag.CommandLine.Parse(nil)
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
