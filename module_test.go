package gogenlicense

import (
	"fmt"
	"os"
)

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
	// Output: github.com/tkw1536/gogenlicense
}
