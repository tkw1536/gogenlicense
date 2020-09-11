// Command gogenlicense finds and formats package licenses for use in go programs.
package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
	"github.com/tkw1536/gogenlicense"
	"github.com/tkw1536/gogenlicense/legal"
)

func main() {
	result, err := gogenlicense.Format(context.Background(), gogenlicense.Options{
		ModulePaths:         importPaths,
		ConfidenceThreshold: confidenceThreshold,

		OutPackage:      outPackage,
		DeclarationName: declarationName,
	})
	if err != nil {
		glog.Fatal(err)
	}

	if outFile != "" {
		err := ioutil.WriteFile(outFile, []byte(result), os.ModePerm)
		if err != nil {
			glog.Fatal(err)
		}
		return
	}

	fmt.Print(result)
}

var importPaths []string
var confidenceThreshold float64 = 0.9
var outFile string = ""
var outPackage string = "legal"
var declarationName string = "Notices"

func init() {
	var legalFlag bool
	flag.BoolVar(&legalFlag, "legal", legalFlag, "Print legal notices and exit")
	defer func() {
		if legalFlag {
			fmt.Println("This executable contains code from several different go packages. ")
			fmt.Println("Some of these packages require licensing information to be made available to the end user. ")
			fmt.Println(legal.Notices)
			os.Exit(0)
		}
	}()

	flag.Float64Var(&confidenceThreshold, "t", confidenceThreshold, "Threshold for the licenseclassifier")
	flag.StringVar(&outPackage, "p", outPackage, "Package of generated go file")
	flag.StringVar(&declarationName, "n", declarationName, "Name of declaration to generate")
	flag.StringVar(&outFile, "d", outFile, "File to write output to. If omitted, use STDOUT. ")
	flag.Parse()
	importPaths = flag.Args()
}
