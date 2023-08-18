# gogenlicense

![CI Status](https://github.com/tkw1536/gogenlicense/workflows/CI/badge.svg)

gogenlicense finds and formats package licenses for use in go programs.
It is intended to comply with license requirements, but has not been vetted by a lawyer.

Concretely this package generates go code that contains a single string variable which includes all legal notices.
It also generates appropriate comments for godoc.

## Quickstart

Install the command using standard go tools:

```bash
go install github.com/tkw1536/gogenlicense/cmd/gogenlicense@latest
```

gogenlicense is most commonly invoked using a 'go:generate' comment like so:

```go
//go:generate gogenlicense -m
```

After adding such a comment, run the [`go generate` command](https://golang.org/pkg/cmd/go/internal/generate/).

Once a source file has been generated, your go code can make use of the generated notices.
For example to add a legal flag that prints notices:

```go
var legalFlag bool
flag.BoolVar(&legalFlag, "legal", legalFlag, "print legal notices and exit")
defer func() {
	if legalFlag {
		fmt.Println("This executable contains code from several different go packages. ")
		fmt.Println("Some of these packages require licensing information to be made available to the end user. ")
		fmt.Println(legal.Notices) // this references the generated constant!
		os.Exit(0)
	}
}()
```

For a more complete example, see the source code of [the gogenlicense command](./cmd/gogenlicense/main.go).
The notices constant is implemented in the [github.com/tkw1536/gogenlicense/legal](./legal/docs.go) subpackage.
The legal flag is implemented in the main package.

## Usage

```
gogenlicense [-legal] [-help] [-o FILE] [-t THRESHOLD] [-p PACKAGE] [-n NAME] [-m] [MODULE [MODULE...]]
-legal
  Print legal notices and exit.
-help
  Print a usage message and exit.
-p PACKAGE
  Package name of go source code file to generate.
  Defaults to the 'GOPACKAGE' environment variable, or to 'notices' if unset.
-n string
  Name of declaration to generate (default 'Notices')
-t THRESHOLD
 Threshold to use for the license classifier, between 0 and 1.
 Defaults to 0.9.
-o filename
  Path to write output to.
  Defaults to appending a suffix '_notices' to the source file pointed to by the 'GOFILE' environment variable.
  If 'GOFILE' is not provided, uses the filename 'notices.go'.
MODULE
  Import path of module to scan for module dependencies.
-m
  automatically find the current 'go module' containing the working directory and include it in MODULE list.
```

## Acknowledgements

This package was inspired by the https://classic.yarnpkg.com/lang/en/docs/cli/licenses/#toc-yarn-licenses-generate-disclaimer command.
Internally the command relies on the excellent https://github.com/google/go-licenses package.
See also the [legal subpackage](./legal/docs_notices.go) for full licenses of used packages.

## Changelog

## Version 1.2.6 (Released [Aug 19 2023](https://github.com/tkw1536/gogenlicense/releases/tag/v1.2.6))

- add 'DO NOT EDIT' comment to template
- update spellchecker template


## Version 1.2.5 (Released [Aug 8 2023](https://github.com/tkw1536/gogenlicense/releases/tag/v1.2.5))

- add 'spellchecker:disable' comment to template
- fix typos in documentation
- update dependencies

## Version 1.2.4 (Released [Aug 15 2022](https://github.com/tkw1536/gogenlicense/releases/tag/v1.2.4))

- properly handle repositories in or under $GOROOT

## Version 1.2.3 (Released [Aug 15 2022](https://github.com/tkw1536/gogenlicense/releases/tag/v1.2.3))

- fix typo in error message

## Version 1.2.2 (Released [Apr 28 2022](https://github.com/tkw1536/gogenlicense/releases/tag/v1.2.2))

- minor dependency updates

## Version 1.2.1 (Released [Nov 18 2021](https://github.com/tkw1536/gogenlicense/releases/tag/v1.2.1))

- Bugfix: Fix using correct templates for godoc and value string

## Version 1.2.0 (Released [Nov 18 2021](https://github.com/tkw1536/gogenlicense/releases/tag/v1.2.0))

- Add lots of documentation.
- Add an '-m' flag to automatically use the current go.module
- Use '$GOPACKAGE' and '$GOFILE' as default parameters
- Update go to `1.17`
- various internal improvements

### Version 1.1.0 (Released [Feb 21 2021](https://github.com/tkw1536/gogenlicense/releases/tag/v1.1.0))

- Bugfix: Do not include modules with '.go' in the license file name

### Version 1.0.0 (Released [Dec 8 2020](https://github.com/tkw1536/gogenlicense/releases/tag/v1.0.0))

- Initial release

<!-- spellchecker: words gogenlicense godoc println gopackage gofile goroot  -->