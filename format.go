package gogenlicense

import (
	"go/format"

	"github.com/pkg/errors"
)

// gofmt formats code in the standard go fmt style.
func gofmt(code string) (string, error) {
	bytes, err := format.Source([]byte(code))
	if err != nil {
		return "", errors.Wrap(err, "Error running gofmt")
	}
	return string(bytes), nil
}
