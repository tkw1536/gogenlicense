# gogenlicense

![CI Status](https://github.com/tkw1536/gogenlicense/workflows/CI/badge.svg)

Gogenlicense finds and formats package licenses for use in go programs.
It is intended to comply with license requirements, but has not been vetted by a lawyer. 

Concretely this package generates go code that contains a single string variable that contains all the required notices. 
It also generates appropriate comments for godoc. 

It is intended to be used by means of the [`go generate`](https://golang.org/pkg/cmd/go/internal/generate/) command with a comment inside the code such as:

```bash
//go:generate gogenlicense -p legal -n Notices -d notices.go github.com/tkw1536/gogenlicense
```

For a concrete example of this, see the 'legal' subpackage, as it actually implements this. 

This package mainly relies on the execellent https://github.com/google/go-licenses package. 
