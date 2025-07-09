package gogenlicense

import (
	"fmt"
	"os"
)

// spellchecker: words gogenlicense

func ExampleFindModulePath() {
	// get the current directory - module containing this example
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// find the import path to the module
	module, err := FindModulePath(dir)
	if err != nil {
		panic(err)
	}
	fmt.Println(module)
	// Output: go.tkw01536.de/gogenlicense
}
