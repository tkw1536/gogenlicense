// Command gogenlicense finds and formats package licenses for use in go programs.
package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/tkw1536/gogenlicense"
	"github.com/tkw1536/gogenlicense/legal"
)

func main() {
	result, err := gogenlicense.Format(context.Background(), gogenlicense.Options{
		ModulePaths:         argImportPaths,
		ConfidenceThreshold: flagConfidenceThreshold,

		OutPackage:      flagOutPackage,
		DeclarationName: flagDeclarationName,
	})
	if err != nil {
		glog.Fatal(err)
	}

	if flagOutFile != "" {
		err := ioutil.WriteFile(flagOutFile, []byte(result), os.ModePerm)
		if err != nil {
			glog.Fatal(err)
		}
		return
	}

	fmt.Print(result)
}

var argImportPaths []string
var flagConfidenceThreshold float64
var flagOutFile string
var flagOutPackage string
var flagDeclarationName string

func init() {
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

	flag.Parse()
	argImportPaths = flag.Args()
}
