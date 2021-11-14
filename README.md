# gogenlicense

![CI Status](https://github.com/tkw1536/gogenlicense/workflows/CI/badge.svg)

Gogenlicense finds and formats package licenses for use in go programs.
It is intended to comply with license requirements, but has not been vetted by a lawyer. 

Concretely this package generates go code that contains a single string variable that contains all the required notices. 
It also generates appropriate comments for godoc. 

It is intended to be used by means of the [`go generate`](https://golang.org/pkg/cmd/go/internal/generate/) command with a comment inside the code such as:

```bash
//go:generate gogenlicense github.com/tkw1536/gogenlicense
```

For a concrete example of this, see the 'legal' subpackage, as it actually implements this. 

This package mainly relies on the excellent https://github.com/google/go-licenses package. 

## Changelog

## Version 1.1.1 (Upcoming)

- Use '$GOPACKAGE' and '$GOFILE' as default parameters

### Version 1.1.0 (Released [Feb 21 2021](https://github.com/tkw1536/gogenlicense/releases/tag/v1.1.0))

- Bugfix: Do not include modules with '.go' in the license file name

### Version 1.0.0 (Released [Dec 8 2020](https://github.com/tkw1536/gogenlicense/releases/tag/v1.0.0))

- Initial release